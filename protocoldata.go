package servicenames

var protocolData = []Protocol{
	{Proto: 0, Keyword: "hopopt"},
	{Proto: 1, Keyword: "icmp"},
	{Proto: 2, Keyword: "igmp"},
	{Proto: 3, Keyword: "ggp"},
	{Proto: 4, Keyword: "ipv4"},
	{Proto: 5, Keyword: "st"},
	{Proto: 6, Keyword: "tcp"},
	{Proto: 7, Keyword: "cbt"},
	{Proto: 8, Keyword: "egp"},
	{Proto: 9, Keyword: "igp"},
	{Proto: 10, Keyword: "bbn-rcc-mon"},
	{Proto: 11, Keyword: "nvp-ii"},
	{Proto: 12, Keyword: "pup"},
	{Proto: 13, Keyword: "argus (deprecated)"},
	{Proto: 14, Keyword: "emcon"},
	{Proto: 15, Keyword: "xnet"},
	{Proto: 16, Keyword: "chaos"},
	{Proto: 17, Keyword: "udp"},
	{Proto: 18, Keyword: "mux"},
	{Proto: 19, Keyword: "dcn-meas"},
	{Proto: 20, Keyword: "hmp"},
	{Proto: 21, Keyword: "prm"},
	{Proto: 22, Keyword: "xns-idp"},
	{Proto: 23, Keyword: "trunk-1"},
	{Proto: 24, Keyword: "trunk-2"},
	{Proto: 25, Keyword: "leaf-1"},
	{Proto: 26, Keyword: "leaf-2"},
	{Proto: 27, Keyword: "rdp"},
	{Proto: 28, Keyword: "irtp"},
	{Proto: 29, Keyword: "iso-tp4"},
	{Proto: 30, Keyword: "netblt"},
	{Proto: 31, Keyword: "mfe-nsp"},
	{Proto: 32, Keyword: "merit-inp"},
	{Proto: 33, Keyword: "dccp"},
	{Proto: 34, Keyword: "3pc"},
	{Proto: 35, Keyword: "idpr"},
	{Proto: 36, Keyword: "xtp"},
	{Proto: 37, Keyword: "ddp"},
	{Proto: 38, Keyword: "idpr-cmtp"},
	{Proto: 39, Keyword: "tp++"},
	{Proto: 40, Keyword: "il"},
	{Proto: 41, Keyword: "ipv6"},
	{Proto: 42, Keyword: "sdrp"},
	{Proto: 43, Keyword: "ipv6-route"},
	{Proto: 44, Keyword: "ipv6-frag"},
	{Proto: 45, Keyword: "idrp"},
	{Proto: 46, Keyword: "rsvp"},
	{Proto: 47, Keyword: "gre"},
	{Proto: 48, Keyword: "dsr"},
	{Proto: 49, Keyword: "bna"},
	{Proto: 50, Keyword: "esp"},
	{Proto: 51, Keyword: "ah"},
	{Proto: 52, Keyword: "i-nlsp"},
	{Proto: 53, Keyword: "swipe (deprecated)"},
	{Proto: 54, Keyword: "narp"},
	{Proto: 55, Keyword: "min-ipv4"},
	{Proto: 56, Keyword: "tlsp"},
	{Proto: 57, Keyword: "skip"},
	{Proto: 58, Keyword: "ipv6-icmp"},
	{Proto: 59, Keyword: "ipv6-nonxt"},
	{Proto: 60, Keyword: "ipv6-opts"},
	{Proto: 61, Keyword: ""},
	{Proto: 62, Keyword: "cftp"},
	{Proto: 63, Keyword: ""},
	{Proto: 64, Keyword: "sat-expak"},
	{Proto: 65, Keyword: "kryptolan"},
	{Proto: 66, Keyword: "rvd"},
	{Proto: 67, Keyword: "ippc"},
	{Proto: 68, Keyword: ""},
	{Proto: 69, Keyword: "sat-mon"},
	{Proto: 70, Keyword: "visa"},
	{Proto: 71, Keyword: "ipcv"},
	{Proto: 72, Keyword: "cpnx"},
	{Proto: 73, Keyword: "cphb"},
	{Proto: 74, Keyword: "wsn"},
	{Proto: 75, Keyword: "pvp"},
	{Proto: 76, Keyword: "br-sat-mon"},
	{Proto: 77, Keyword: "sun-nd"},
	{Proto: 78, Keyword: "wb-mon"},
	{Proto: 79, Keyword: "wb-expak"},
	{Proto: 80, Keyword: "iso-ip"},
	{Proto: 81, Keyword: "vmtp"},
	{Proto: 82, Keyword: "secure-vmtp"},
	{Proto: 83, Keyword: "vines"},
	{Proto: 84, Keyword: "iptm"},
	{Proto: 85, Keyword: "nsfnet-igp"},
	{Proto: 86, Keyword: "dgp"},
	{Proto: 87, Keyword: "tcf"},
	{Proto: 88, Keyword: "eigrp"},
	{Proto: 89, Keyword: "ospfigp"},
	{Proto: 90, Keyword: "sprite-rpc"},
	{Proto: 91, Keyword: "larp"},
	{Proto: 92, Keyword: "mtp"},
	{Proto: 93, Keyword: "ax.25"},
	{Proto: 94, Keyword: "ipip"},
	{Proto: 95, Keyword: "micp (deprecated)"},
	{Proto: 96, Keyword: "scc-sp"},
	{Proto: 97, Keyword: "etherip"},
	{Proto: 98, Keyword: "encap"},
	{Proto: 99, Keyword: ""},
	{Proto: 100, Keyword: "gmtp"},
	{Proto: 101, Keyword: "ifmp"},
	{Proto: 102, Keyword: "pnni"},
	{Proto: 103, Keyword: "pim"},
	{Proto: 104, Keyword: "aris"},
	{Proto: 105, Keyword: "scps"},
	{Proto: 106, Keyword: "qnx"},
	{Proto: 107, Keyword: "a/n"},
	{Proto: 108, Keyword: "ipcomp"},
	{Proto: 109, Keyword: "snp"},
	{Proto: 110, Keyword: "compaq-peer"},
	{Proto: 111, Keyword: "ipx-in-ip"},
	{Proto: 112, Keyword: "vrrp"},
	{Proto: 113, Keyword: "pgm"},
	{Proto: 114, Keyword: ""},
	{Proto: 115, Keyword: "l2tp"},
	{Proto: 116, Keyword: "ddx"},
	{Proto: 117, Keyword: "iatp"},
	{Proto: 118, Keyword: "stp"},
	{Proto: 119, Keyword: "srp"},
	{Proto: 120, Keyword: "uti"},
	{Proto: 121, Keyword: "smp"},
	{Proto: 122, Keyword: "sm (deprecated)"},
	{Proto: 123, Keyword: "ptp"},
	{Proto: 124, Keyword: "isis over ipv4"},
	{Proto: 125, Keyword: "fire"},
	{Proto: 126, Keyword: "crtp"},
	{Proto: 127, Keyword: "crudp"},
	{Proto: 128, Keyword: "sscopmce"},
	{Proto: 129, Keyword: "iplt"},
	{Proto: 130, Keyword: "sps"},
	{Proto: 131, Keyword: "pipe"},
	{Proto: 132, Keyword: "sctp"},
	{Proto: 133, Keyword: "fc"},
	{Proto: 134, Keyword: "rsvp-e2e-ignore"},
	{Proto: 135, Keyword: "mobility header"},
	{Proto: 136, Keyword: "udplite"},
	{Proto: 137, Keyword: "mpls-in-ip"},
	{Proto: 138, Keyword: "manet"},
	{Proto: 139, Keyword: "hip"},
	{Proto: 140, Keyword: "shim6"},
	{Proto: 141, Keyword: "wesp"},
	{Proto: 142, Keyword: "rohc"},
	{Proto: 143, Keyword: "ethernet"},
	{Proto: 144, Keyword: "aggfrag"},
	{Proto: 145, Keyword: "nsh"},
	{Proto: 253, Keyword: ""},
	{Proto: 254, Keyword: ""},
	{Proto: 255, Keyword: "reserved"},
}
