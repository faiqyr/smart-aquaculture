package infrastructure

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase-admin.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Peringatan: Gagal inisialisasi Firebase (pastikan firebase-admin.json ada): %v\n", err)
		return
	}
	FirebaseApp = app
	log.Println("Firebase Admin SDK berhasil diinisialisasi")
}
