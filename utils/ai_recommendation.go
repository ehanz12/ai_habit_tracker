package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

)

func CallGroq(prompt string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")

	url := "https://api.groq.com/openai/v1/chat/completions"

	reqBody := map[string]interface{}{
		"model": "llama-3.1-8b-instant",
		"messages": []map[string]string{
			{
				"role": "user",
				"content": prompt,
			},
		},
		"temperature": 0.7,
	}

	jsonData, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// 🔥 SAFE PARSING
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("invalid response: %v", result)
	}

	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	content := message["content"].(string)

	return content, nil
}


func GenerateAIInsight(habitName string, completed, missed, streak int) string {
	prompt := fmt.Sprintf(`
Kamu adalah AI habit coach yang santai dan memotivasi.

Data:
- Habit: %s
- Streak: %d hari
- Berhasil: %d hari
- Gagal: %d hari

Buat insight singkat (maks 2 kalimat), pakai emoji.
dan berikan tips sederhana untuk meningkatkan konsistensi habit ini.
atau berikan recommendasi untuk habit baru jika habit ini sudah tidak efektif lagi.
`, habitName, streak, completed, missed)

	text, err := CallGroq(prompt)
	if err != nil {
		return fallbackInsight(completed, missed, streak)
	}

	return text
}


func fallbackInsight(completed, missed, streak int) string {
	if streak >= 5 {
		return "🔥 Kamu konsisten banget, lanjutkan!"
	}
	if completed > missed {
		return "💪 Progress bagus, tinggal sedikit lagi jadi habit kuat!"
	}
	return "😅 Kamu mulai kurang konsisten, coba ubah waktu habit"
}