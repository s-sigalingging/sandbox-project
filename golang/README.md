This code can be used to list Google Compute Engine Instances
The output of this code is Instance Name, Instance Status, Network, and Machine Type. Though, for the network and machine type still need some improvement to produce just the IP Adress and the machine type.
To run the code we have to put the service account key first in the sa.json file and make sure the service account have enough permission to list compute engine instances in your project. Don't forget to change your project_id and zone in main.go
After you fulfill all of the steps, execute the code by run:
go run main.go