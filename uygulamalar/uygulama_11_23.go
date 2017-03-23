// uygulama_11_23
package main

import gexpect "github.com/ThomasRooney/gexpect"
import "log"

func main() {
	log.Printf("FTP bağlantısı açılıyor... ")
	ftp, err := gexpect.Spawn("ftp ftp.linux.org.tr")
	if err != nil {
		panic(err)
	}
	log.Printf("Bağlantı kuruldu, isim sorulması bekleniyor ")
	ftp.Expect("Name")
	ftp.SendLine("anonymous")
	log.Printf("İsim gönderildi")
	ftp.Expect("Password:")
	ftp.SendLine("eposta@example.com")
	log.Printf("Parola gönderildi")
	ftp.Expect("ftp> ")
	ftp.SendLine("cd /ubuntu-releases/yakkety/")
	log.Printf("Klasör değiştirildi")
	ftp.SendLine("get MD5SUMS")
	log.Printf("Dosya alındı")
	ftp.Expect("ftp> ")
	ftp.SendLine("bye")
	ftp.Expect("221")
	log.Printf("Oturum tamamlandı.\n")
}
