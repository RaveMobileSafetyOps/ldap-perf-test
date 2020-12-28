package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"math/rand"
	"time"

	goldap "github.com/jbuchbinder/goldap"
)

var (
	LdapHost     = flag.String("ldaphost", "127.0.0.1", "LDAP host")
	LdapPort     = flag.Int("ldapport", 389, "LDAP port")
	LdapSsl      = flag.Bool("ldapssl", false, "LDAP SSL enabled")
	LdapBaseDn   = flag.String("ldapbase", "", "LDAP base DN")
	LdapBindDn   = flag.String("ldapbind", "", "LDAP bind DN")
	LdapPassword = flag.String("ldappass", "", "LDAP password")
	LdapSearch   = flag.String("ldapsearch", "", "LDAP search string")
	Threads      = flag.Int("threads", 1, "Test threads")
	Delay        = flag.Int("delay", 5000, "Delay randomization")
	LaxSsl       = flag.Bool("laxssl", false, "Use lax SSL restrictions")
)

func main() {
	flag.Parse()

	cpool := make([]*goldap.Conn, *Threads+1)
	for i := 1; i <= *Threads; i++ {
		//fmt.Printf("Launching thread %d\n", i)
		testBind(i, cpool[i])
	}
}

func testBind(threadId int, l *goldap.Conn) {
	fmt.Printf("Test [thread %d]: starting...\n", threadId)
	delay := rand.Intn(*Delay)
	fmt.Printf("Test [thread %d]: waiting for %d ms\n", threadId, delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	t := NewTimer()
	var err *goldap.Error
	if *LdapSsl {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: *LaxSsl,
		}
		l, err = goldap.DialSSLConfig("tcp", fmt.Sprintf("%s:%d", *LdapHost, *LdapPort), tlsConfig)
	} else {
		l, err = goldap.Dial("tcp", fmt.Sprintf("%s:%d", *LdapHost, *LdapPort))
	}
	t.End()
	fmt.Printf("Dial [thread %d] - %s\n", threadId, t.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//defer l.Close()

	t = NewTimer()
	err = l.Bind(*LdapBindDn, *LdapPassword)
	t.End()
	fmt.Printf("Bind [thread %d] - %s\n", threadId, t.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *LdapSearch != "" {
		search := goldap.NewSearchRequest(
			*LdapBaseDn,
			goldap.ScopeWholeSubtree, goldap.DerefAlways,
			0, 0, false,
			*LdapSearch,
			[]string{
				"cn",
				"description"},
			nil)
		t = NewTimer()
		_, err := l.Search(search)
		t.End()
		fmt.Printf("Search [thread %d] - %s\n", threadId, t.String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
