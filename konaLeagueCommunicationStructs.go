package espnffgo

type KonaLeagueCommunication struct {
	Communication struct {
		Topics []struct {
			Author       string `json:"author"`
			CreationInfo struct {
				ClientAddress any    `json:"clientAddress"`
				Date          int64  `json:"date"`
				Platform      string `json:"platform"`
				Source        string `json:"source"`
			} `json:"creationInfo"`
			Date           int64  `json:"date"`
			DateEdited     int64  `json:"dateEdited"`
			ID             string `json:"id"`
			IsDeleted      bool   `json:"isDeleted"`
			IsEdited       bool   `json:"isEdited"`
			LastUpdateInfo struct {
				ClientAddress any    `json:"clientAddress"`
				Date          int64  `json:"date"`
				Platform      string `json:"platform"`
				Source        string `json:"source"`
			} `json:"lastUpdateInfo"`
			Messages []struct {
				Author       string `json:"author"`
				CreationInfo struct {
					ClientAddress any    `json:"clientAddress"`
					Date          int64  `json:"date"`
					Platform      string `json:"platform"`
					Source        string `json:"source"`
				} `json:"creationInfo"`
				Date                int64  `json:"date"`
				For                 int    `json:"for"`
				From                int    `json:"from"`
				ID                  string `json:"id"`
				IsAlternativeFormat bool   `json:"isAlternativeFormat"`
				IsDeleted           bool   `json:"isDeleted"`
				IsEdited            bool   `json:"isEdited"`
				MessageTypeID       int    `json:"messageTypeId"`
				TargetID            int64  `json:"targetId"`
				To                  int    `json:"to"`
				TopicID             string `json:"topicId"`
			} `json:"messages"`
			TargetID          int64  `json:"targetId"`
			Subject           string `json:"subject"`
			TotalMessageCount int64  `json:"totalMessageCount"`
			Type              string `json:"type"`
		} `json:"topics"`
	} `json:"communication"`
}
