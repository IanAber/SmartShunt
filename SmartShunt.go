package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"go.bug.st/serial"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var (
	WebPort  string
	webFiles string
	Params   ParamsType
	s        serial.Port
)

func defaultPage(w http.ResponseWriter, r *http.Request) {
	if buffer, err := Params.getJSON(); err != nil {
		log.Println(err)
	} else {
		if _, err := w.Write(buffer); err != nil {
			log.Println(err)
		}
	}
}

func setUpWebSite() {
	router := mux.NewRouter()

	router.HandleFunc("/", defaultPage).Methods("GET")
	port := fmt.Sprintf(":%s", WebPort)
	log.Fatal(http.ListenAndServe(port, router))
}

func init() {
	flag.StringVar(&WebPort, "WebPort", "28000", "Web port")
	flag.StringVar(&webFiles, "webFiles", "/SmartShunt/web", "Path to the WEB files location")
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Starting the WEB site.")
	log.Println("Starting the serial port")
	if s = ConnectSerial(); s == nil {
		log.Println("Failed to strat, trying again...")
		s = RestartSerial()
		if s == nil {
			log.Fatal("Failed to open the serial port on startup.")
		}
	}
	log.Println("Serial port is running")
	go setUpWebSite()
}

func ConnectSerial() serial.Port {
	mode := &serial.Mode{
		BaudRate: 19200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := port.SetReadTimeout(time.Millisecond * 150); err != nil {
		log.Println(err)
	}
	log.Println("Serial port connected")
	return port
}

func RestartSerial() serial.Port {
	log.Println("Restarting serial port")
	if s != nil {
		log.Println("Closing ttyUSB0")
		if err := s.Close(); err != nil {
			log.Println(err)
		}
		s = nil
	}
	log.Println("Port closed")
	time.Sleep(time.Second * 2)
	log.Println("Running usbreset command")
	// Reset the USB device by using the usbreset command line program
	cmd := exec.Command("usbreset", "1a86:7523")
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return nil
	} else {
		log.Println(string(stdout))
	}
	log.Println("Reset complete, connecting")
	time.Sleep(time.Second * 2)
	return ConnectSerial()
}

func main() {
	log.Println("Starting main.")

	paramBuf := make([]byte, 128)
	idx := 0
	timeout := 0
	for {
		if s == nil {
			s = RestartSerial()
			if s == nil {
				log.Println("Failed to connect")
			}
		}
		if s != nil {
			buf := make([]byte, 1)
			_, err := s.Read(buf)
			timeout++
			if err != nil && err != io.EOF {
				s = RestartSerial()
				timeout = 0
			}
			paramBuf[idx] = buf[0]
			idx++
			timeout = 0
			if buf[0] == 0x0A {
				//				paramBuf[idx-2] = 0
				paramBuf[idx-1] = 0
				Params.setValues(string(paramBuf[0 : idx-2]))
				idx = 0
			}
			if idx > len(paramBuf) {
				idx = 0
			}
			if timeout > 10 {
				log.Println("Timed out...")
				s = RestartSerial()
				timeout = 0
				idx = 0
			}
		}
	}
}
