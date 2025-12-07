package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {
	flagLen := 13

	flag := []byte(strings.Repeat(string(rune(33)), flagLen))

	for i := range flagLen {
		for range 93 {
			flag[i]++

			cmd := exec.Command("./challenge-prod")
			stdin, err := cmd.StdinPipe()
			if err != nil {
				panic(err)
			}

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				panic(err)
			}
			if err := cmd.Start(); err != nil {
				panic(err)
			}

			if _, err := io.WriteString(stdin, string(flag)+"\n"); err != nil {
				panic(err)
			}
			if err := stdin.Close(); err != nil {
				panic(err)
			}

			output, err := io.ReadAll(stdout)
			if err != nil {
				panic(err)
			}

			if !strings.HasSuffix(string(output), fmt.Sprintf("Wrong at index %d\n", i)) {
				fmt.Println(string(output))
				if strings.Contains(string(output), "Correct") {
					fmt.Println("flag: ", string(flag))
				}
				break
			}

			cmd.Wait()
		}
	}
}

// 以下のコードで長さが13と分かった
//func main() {
//	for i := range 32 {
//		flagLen := i + 1
//		fmt.Printf("flag_len: %d\n", flagLen)
//		cmd := exec.Command("./challenge-prod")
//		stdin, err := cmd.StdinPipe()
//		if err != nil {
//			panic(err)
//		}
//
//		stdout, err := cmd.StdoutPipe()
//		if err != nil {
//			panic(err)
//		}
//		if err := cmd.Start(); err != nil {
//			panic(err)
//		}
//
//		in := strings.Repeat("a", flagLen)
//		if _, err := io.WriteString(stdin, in+"\n"); err != nil {
//			panic(err)
//		}
//		if err := stdin.Close(); err != nil {
//			panic(err)
//		}
//
//		output, err := io.ReadAll(stdout)
//		if err != nil {
//			log.Print(err)
//		}
//
//		fmt.Println(string(output))
//
//		if err := cmd.Wait(); err != nil {
//			log.Print(err)
//		}
//	}
//}
