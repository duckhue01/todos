package db

import (
	"fmt"
	"math/rand"
	"path"
	"time"
)

func (db *jsonDB) ReadMedievalMusic() string {
	rand.Seed(time.Now().UnixNano())
	return path.Join(db.path, fmt.Sprintf("musics/medieval/%d.mp3", rand.Intn(8)))
}

func (db *jsonDB) ReadEpicMusic() string {
	rand.Seed(time.Now().UnixNano())
	return path.Join(db.path, fmt.Sprintf("musics/epic/%d.mp3", rand.Intn(8)))
}

func (db *jsonDB) ReadPianoMusic() string {
	rand.Seed(time.Now().UnixNano())
	return path.Join(db.path, fmt.Sprintf("musics/piano/%d.mp3", rand.Intn(8)))
}

func (db *jsonDB) ReadChillMusic() string {
	rand.Seed(time.Now().UnixNano())
	return path.Join(db.path, fmt.Sprintf("musics/chill/%d.mp3", rand.Intn(8)))
}
