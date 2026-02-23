package main

func main() {
	bc := newBlockChain()
	defer bc.db.Close()

	cli := CLI{bc: bc}
	cli.Run()
}
