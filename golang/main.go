package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// service adalah variable yang di desain untuk memanggil sa.json yang digunakan sebagai autentikasi
	service, err := compute.NewService(ctx, option.WithCredentialsFile("sa.json"))
	if err != nil {
		log.Fatalf("Gagal Connect %v", err)
	}
	// fmt.Println(service)

	// instance adalah variable yang digunakan untuk call api instance list
	instances, err := service.Instances.List("project_id", "zone").Do()
	if err != nil {
		log.Fatalf("Gagal mengambil instance list %v", err)
	}
	// fmt.Println(instance)
	for _, instance := range instances.Items {
		fmt.Printf("Nama: %s\n", instance.Name)
		fmt.Printf("Status: %s\n", instance.Status)
		fmt.Printf("Network: %s\n", instance.NetworkInterfaces)
		fmt.Printf("Type: %s\n", instance.MachineType)
		fmt.Println("================================")

	}

}
