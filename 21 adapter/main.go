package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере

type Player struct{}

func (p *Player) Play(f File) {
	f.LoadFile()
	fmt.Println("MUSIC!")
}

type File interface {
	LoadFile()
}

type MP3 struct{}

func (mp3 *MP3) LoadFile() {
	fmt.Println("mp3 loaded")
}

type VAW struct{}

func (vaw *VAW) LoadFileVaw() {
	fmt.Println("vaw loaded")
}

type VAWadapterToMp3 struct {
	file *VAW
}

func (adapter *VAWadapterToMp3) LoadFile() {
	fmt.Println("vaw converted to mp3")
	adapter.file.LoadFileVaw()
}

func main() {
	player := Player{}
	//Playing mp3
	mp3 := MP3{}
	player.Play(&mp3)

	//Convert
	vaw := VAW{}
	vawApater := VAWadapterToMp3{file: &vaw}

	//Playing vaw
	player.Play(&vawApater)
}
