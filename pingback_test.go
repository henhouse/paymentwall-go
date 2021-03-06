package paymentwall

import "testing"

func TestPingback_IsIPValid(t *testing.T) {
	var tests = []struct {
		ip      string
		isValid bool
	}{
		// Invalid IPs.
		{"1.2.3.4", false},
		{"127.0.0.1", false},
		{"192.168.1.1", false},
		{"216.127.72.1", false},
		{"216.127.73.2", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", false}, // IPv6 not supported.

		// All valid, within-range IPs.
		{"216.127.71.0", true},
		{"216.127.71.1", true},
		{"216.127.71.2", true},
		{"216.127.71.3", true},
		{"216.127.71.4", true},
		{"216.127.71.5", true},
		{"216.127.71.6", true},
		{"216.127.71.7", true},
		{"216.127.71.8", true},
		{"216.127.71.9", true},
		{"216.127.71.10", true},
		{"216.127.71.11", true},
		{"216.127.71.12", true},
		{"216.127.71.13", true},
		{"216.127.71.14", true},
		{"216.127.71.15", true},
		{"216.127.71.16", true},
		{"216.127.71.17", true},
		{"216.127.71.18", true},
		{"216.127.71.19", true},
		{"216.127.71.20", true},
		{"216.127.71.21", true},
		{"216.127.71.22", true},
		{"216.127.71.23", true},
		{"216.127.71.24", true},
		{"216.127.71.25", true},
		{"216.127.71.26", true},
		{"216.127.71.27", true},
		{"216.127.71.28", true},
		{"216.127.71.29", true},
		{"216.127.71.30", true},
		{"216.127.71.31", true},
		{"216.127.71.32", true},
		{"216.127.71.33", true},
		{"216.127.71.34", true},
		{"216.127.71.35", true},
		{"216.127.71.36", true},
		{"216.127.71.37", true},
		{"216.127.71.38", true},
		{"216.127.71.39", true},
		{"216.127.71.40", true},
		{"216.127.71.41", true},
		{"216.127.71.42", true},
		{"216.127.71.43", true},
		{"216.127.71.44", true},
		{"216.127.71.45", true},
		{"216.127.71.46", true},
		{"216.127.71.47", true},
		{"216.127.71.48", true},
		{"216.127.71.49", true},
		{"216.127.71.50", true},
		{"216.127.71.51", true},
		{"216.127.71.52", true},
		{"216.127.71.53", true},
		{"216.127.71.54", true},
		{"216.127.71.55", true},
		{"216.127.71.56", true},
		{"216.127.71.57", true},
		{"216.127.71.58", true},
		{"216.127.71.59", true},
		{"216.127.71.60", true},
		{"216.127.71.61", true},
		{"216.127.71.62", true},
		{"216.127.71.63", true},
		{"216.127.71.64", true},
		{"216.127.71.65", true},
		{"216.127.71.66", true},
		{"216.127.71.67", true},
		{"216.127.71.68", true},
		{"216.127.71.69", true},
		{"216.127.71.70", true},
		{"216.127.71.71", true},
		{"216.127.71.72", true},
		{"216.127.71.73", true},
		{"216.127.71.74", true},
		{"216.127.71.75", true},
		{"216.127.71.76", true},
		{"216.127.71.77", true},
		{"216.127.71.78", true},
		{"216.127.71.79", true},
		{"216.127.71.80", true},
		{"216.127.71.81", true},
		{"216.127.71.82", true},
		{"216.127.71.83", true},
		{"216.127.71.84", true},
		{"216.127.71.85", true},
		{"216.127.71.86", true},
		{"216.127.71.87", true},
		{"216.127.71.88", true},
		{"216.127.71.89", true},
		{"216.127.71.90", true},
		{"216.127.71.91", true},
		{"216.127.71.92", true},
		{"216.127.71.93", true},
		{"216.127.71.94", true},
		{"216.127.71.95", true},
		{"216.127.71.96", true},
		{"216.127.71.97", true},
		{"216.127.71.98", true},
		{"216.127.71.99", true},
		{"216.127.71.100", true},
		{"216.127.71.101", true},
		{"216.127.71.102", true},
		{"216.127.71.103", true},
		{"216.127.71.104", true},
		{"216.127.71.105", true},
		{"216.127.71.106", true},
		{"216.127.71.107", true},
		{"216.127.71.108", true},
		{"216.127.71.109", true},
		{"216.127.71.110", true},
		{"216.127.71.111", true},
		{"216.127.71.112", true},
		{"216.127.71.113", true},
		{"216.127.71.114", true},
		{"216.127.71.115", true},
		{"216.127.71.116", true},
		{"216.127.71.117", true},
		{"216.127.71.118", true},
		{"216.127.71.119", true},
		{"216.127.71.120", true},
		{"216.127.71.121", true},
		{"216.127.71.122", true},
		{"216.127.71.123", true},
		{"216.127.71.124", true},
		{"216.127.71.125", true},
		{"216.127.71.126", true},
		{"216.127.71.127", true},
		{"216.127.71.128", true},
		{"216.127.71.129", true},
		{"216.127.71.130", true},
		{"216.127.71.131", true},
		{"216.127.71.132", true},
		{"216.127.71.133", true},
		{"216.127.71.134", true},
		{"216.127.71.135", true},
		{"216.127.71.136", true},
		{"216.127.71.137", true},
		{"216.127.71.138", true},
		{"216.127.71.139", true},
		{"216.127.71.140", true},
		{"216.127.71.141", true},
		{"216.127.71.142", true},
		{"216.127.71.143", true},
		{"216.127.71.144", true},
		{"216.127.71.145", true},
		{"216.127.71.146", true},
		{"216.127.71.147", true},
		{"216.127.71.148", true},
		{"216.127.71.149", true},
		{"216.127.71.150", true},
		{"216.127.71.151", true},
		{"216.127.71.152", true},
		{"216.127.71.153", true},
		{"216.127.71.154", true},
		{"216.127.71.155", true},
		{"216.127.71.156", true},
		{"216.127.71.157", true},
		{"216.127.71.158", true},
		{"216.127.71.159", true},
		{"216.127.71.160", true},
		{"216.127.71.161", true},
		{"216.127.71.162", true},
		{"216.127.71.163", true},
		{"216.127.71.164", true},
		{"216.127.71.165", true},
		{"216.127.71.166", true},
		{"216.127.71.167", true},
		{"216.127.71.168", true},
		{"216.127.71.169", true},
		{"216.127.71.170", true},
		{"216.127.71.171", true},
		{"216.127.71.172", true},
		{"216.127.71.173", true},
		{"216.127.71.174", true},
		{"216.127.71.175", true},
		{"216.127.71.176", true},
		{"216.127.71.177", true},
		{"216.127.71.178", true},
		{"216.127.71.179", true},
		{"216.127.71.180", true},
		{"216.127.71.181", true},
		{"216.127.71.182", true},
		{"216.127.71.183", true},
		{"216.127.71.184", true},
		{"216.127.71.185", true},
		{"216.127.71.186", true},
		{"216.127.71.187", true},
		{"216.127.71.188", true},
		{"216.127.71.189", true},
		{"216.127.71.190", true},
		{"216.127.71.191", true},
		{"216.127.71.192", true},
		{"216.127.71.193", true},
		{"216.127.71.194", true},
		{"216.127.71.195", true},
		{"216.127.71.196", true},
		{"216.127.71.197", true},
		{"216.127.71.198", true},
		{"216.127.71.199", true},
		{"216.127.71.200", true},
		{"216.127.71.201", true},
		{"216.127.71.202", true},
		{"216.127.71.203", true},
		{"216.127.71.204", true},
		{"216.127.71.205", true},
		{"216.127.71.206", true},
		{"216.127.71.207", true},
		{"216.127.71.208", true},
		{"216.127.71.209", true},
		{"216.127.71.210", true},
		{"216.127.71.211", true},
		{"216.127.71.212", true},
		{"216.127.71.213", true},
		{"216.127.71.214", true},
		{"216.127.71.215", true},
		{"216.127.71.216", true},
		{"216.127.71.217", true},
		{"216.127.71.218", true},
		{"216.127.71.219", true},
		{"216.127.71.220", true},
		{"216.127.71.221", true},
		{"216.127.71.222", true},
		{"216.127.71.223", true},
		{"216.127.71.224", true},
		{"216.127.71.225", true},
		{"216.127.71.226", true},
		{"216.127.71.227", true},
		{"216.127.71.228", true},
		{"216.127.71.229", true},
		{"216.127.71.230", true},
		{"216.127.71.231", true},
		{"216.127.71.232", true},
		{"216.127.71.233", true},
		{"216.127.71.234", true},
		{"216.127.71.235", true},
		{"216.127.71.236", true},
		{"216.127.71.237", true},
		{"216.127.71.238", true},
		{"216.127.71.239", true},
		{"216.127.71.240", true},
		{"216.127.71.241", true},
		{"216.127.71.242", true},
		{"216.127.71.243", true},
		{"216.127.71.244", true},
		{"216.127.71.245", true},
		{"216.127.71.246", true},
		{"216.127.71.247", true},
		{"216.127.71.248", true},
		{"216.127.71.249", true},
		{"216.127.71.250", true},
		{"216.127.71.251", true},
		{"216.127.71.252", true},
		{"216.127.71.253", true},
		{"216.127.71.254", true},
		{"216.127.71.255", true},
	}

	for _, test := range tests {
		p := &Pingback{ip: test.ip}

		if !p.IsIPValid() && test.isValid {
			t.Errorf("IP was not within range of whitelist, got: %v", test.ip)
		}
	}
}
