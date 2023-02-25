package randomorg

func (r *Random) parseSignature(result map[string]interface{}) string {
	sig, _ := result["signature"]
	if sig != nil {
		return sig.(string)
	}
	return ""
}
