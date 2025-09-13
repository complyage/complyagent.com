package calls

type AgentStruct struct {
	URL   string
	Port  string
	Route string
}

var (
	AGENT_OCR = AgentStruct{
		URL:   "http://localhost",
		Port:  "50051",
		Route: "/ocr",
	}
	AGENT_FACE_DETECT = AgentStruct{
		URL:   "http://localhost",
		Port:  "50052",
		Route: "/detect",
	}
	AGENT_FACE_COMPARE = AgentStruct{
		URL:   "http://localhost",
		Port:  "50052",
		Route: "/compare",
	}
	AGENT_NSFW = AgentStruct{
		URL:   "http://localhost",
		Port:  "50053",
		Route: "/detect",
	}
	AGENT_PHI3 = AgentStruct{
		URL:   "http://localhost",
		Port:  "50053",
		Route: "/api/generate",
	}
	AGENT_DOB = AgentStruct{
		URL:   "http://localhost",
		Port:  "50054",
		Route: "/dob",
	}
)

func (a AgentStruct) GenerateURL() string {
	url := a.URL
	if a.Port != "" {
		url += ":" + a.Port
	}
	if a.Route != "" {
		url += a.Route
	}
	return url
}
