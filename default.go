package k0s

const (
	DEFAULT_HUB_ADDRESS = "https://k0s.up.railway.app"
	DEFAULT_UI_ADDRESS  = "https://k0s.vercel.app"

	REDIR_PROXY_PORT  = ":1081"
	SOCKS5_PROXY_PORT = ":1080"
	DOH_PROXY_PORT    = ":53"
	HUB_PORT          = ":8000"

	MAX_WS_MESSAGE = 8 * 1024 * 1024
)
