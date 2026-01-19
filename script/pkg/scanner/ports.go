package scanner

var PopularPorts = []int{
	// ftp / ssh / telnet
	21, 22, 23,
	// mail
	25, 110, 143, 465, 587, 993, 995,
	// dns
	53,
	// web (Hypertext Transfer Protocol Secure)
	80, 81, 88, 443, 444, 591, 8000, 8008, 8080, 8081, 8088, 8443, 8888,
	// proxy / vpn
	1080, 1194, 3128,
	// databases
	1433, // MSSQL
	1521, // Oracle
	2049, // NFS
	2375, // Docker
	2376,
	2483, 2484, // Oracle SSL
	27017, 27018, // MongoDB
	3306, // MySQL
	5432, // PostgreSQL
	6379, // Redis
	6380,
	9042, // Cassandra
	9200, // Elasticsearch
	9300,
	// remote access
	3389,       // RDP
	5900, 5901, // VNC
	5800,
	// File sharing
	111,      // RPC
	139, 445, // SMB
	// Monitoring / Admin
	10000,      // Webmin
	5601,       // Kibana
	9000, 9001, // Sonar / Admin panels
	// messaging / queues
	5672,  // RabbitMQ
	61616, // ActiveMQ
	// game / misc
	25565, // Minecraft
	11211, // Memcached
}

func AllPorts() []int {
	ports := make([]int, maxPorts)
	for i := 0; i < maxPorts; i++ {
		ports[i] = i + 1
	}
	return ports
}
