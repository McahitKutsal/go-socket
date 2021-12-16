package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Start(server net.Listener) {
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	deadConns := make(chan net.Conn)
	messages := make(chan string)
	i := 0
	j := 1
	//Gelen bağlantıların kabulu
	go func() {
		for {
			conn, err := server.Accept()
			fmt.Println("Client", j, ": Bağlantı kabul edildi")
			j++
			if err != nil {
				log.Println(err.Error())
			}
			conns <- conn
		}
	}()
	for {
		//Gelen bağlantıların okunması
		select {
		//Bir bağlantı aldığımızda onun mesajlarını okumaya başlarız
		case conn := <-conns:
			aconns[conn] = i
			i++
			go func(conn net.Conn, i int) {
				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('@')
					if err != nil {
						break
					}
					if len(m) > 0 {
						messages <- fmt.Sprintf("Client %v: %v", i, m)
					}
				}
				deadConns <- conn
			}(conn, i)
		case msg := <-messages:
			//Mesaj kanalında bir mesaj var ise bunu tüm bağlantılara yayınlıyoruz
			for conn := range aconns {
				conn.Write([]byte(msg))
			}
		case dconn := <-deadConns:
			log.Printf("Client: %v ayrıldı...\n", aconns[dconn])
			delete(aconns, dconn)
		}
	}
}
