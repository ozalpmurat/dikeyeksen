// uygulama_11_24
package main

import gexpect "github.com/ThomasRooney/gexpect"
import "log"

func main() {
	log.Printf("SSH bağlantısı açılıyor... ")
	ssh, err := gexpect.Spawn("ssh murat@localhost")
	if err != nil {
		panic(err)
	}
	ssh.Expect("password:")
	ssh.SendLine("pAr01A")
	ssh.Expect("$")
	ssh.SendLine("passwd")
	ssh.Expect("current")
	ssh.SendLine("pAr01A")
	ssh.Expect("Enter new")
	ssh.SendLine("Yeni_ParoLA")
	ssh.Expect("Retype")
	ssh.SendLine("Yeni_ParoLA")
	ssh.Expect("$")
	ssh.SendLine("logout")
	log.Printf("Parola değiştirildi.\n")
}
