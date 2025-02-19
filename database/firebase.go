package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseClients struct {
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
}

// melakukan inisialisasi menggunakan firebase firestore
func InitFirebase() *FirebaseClients {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	// inisialisasi firebase
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("Eror melakukan inisialisasi", err)
	}

	// inisialisasi auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Error melakukan inisialisasi client auth", err)
	}

	// inisialisasi Firestore client
	firestoreClient, err := firestore.NewClient(context.Background(), "mobile-aplication-34f11", opt)
	if err != nil {
		log.Fatal("Error melakukan inisialisasi client firestore", err)
	}

	return &FirebaseClients{
		AuthClient:      authClient,
		FirestoreClient: firestoreClient,
	}
}
