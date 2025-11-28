package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SeqConfig struct {
	URL     string
	APIKey  string
	Enabled bool
}

var seqCfg SeqConfig

func InitSeq(cfg SeqConfig) {
	seqCfg = cfg
	log.Printf("[INFO] %d | Seq logging enabled: %v", time.Now().Unix(), cfg.Enabled)
}

func Info(msg string) {
	writeLog("INFO", msg, nil)
}

func Error(msg string, err error) {
	writeLog("ERROR", msg, err)
}

func writeLog(level, msg string, err error) {
	now := time.Now().Unix()

	if err != nil {
		log.Printf("[%s] %d | %s: %v", level, now, msg, err)
	} else {
		log.Printf("[%s] %d | %s", level, now, msg)
	}

	if seqCfg.Enabled {
		go sendToSeq(level, msg, err)
	}
}

func sendToSeq(level, msg string, err error) {
	body := map[string]interface{}{
		"@l":      level,
		"@m":      msg,
		"unix_ts": time.Now().Unix(),
		"@t":      time.Now().Format(time.RFC3339),
	}
	if err != nil {
		body["@x"] = err.Error()
	}

	j, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", seqCfg.URL, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if seqCfg.APIKey != "" {
		req.Header.Set("X-Seq-ApiKey", seqCfg.APIKey)
	}
	_, _ = http.DefaultClient.Do(req) // ошибки игнорируем
}
