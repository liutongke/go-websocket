package websocket

// 获取本地分组的成员
func (h *Hub) GetGroupClientList(groupId int) []*Client {
	h.GroupLock.RLock()
	defer h.GroupLock.RUnlock()
	return h.Groups[groupId]
}

// 添加到本地分组
func (h *Hub) addClient2Group(groupId int, client *Client) {
	h.GroupLock.Lock()
	defer h.GroupLock.Unlock()
	h.Groups[groupId] = append(h.Groups[groupId], client)
}

// 删除分组里的客户端
func (h *Hub) delGroupClient(groupId int, client *Client) {
	h.GroupLock.Lock()
	defer h.GroupLock.Unlock()

	for index, v := range h.Groups[groupId] {
		if v.UserId == client.UserId {
			h.Groups[groupId] = append(h.Groups[groupId][:index], h.Groups[groupId][index+1:]...)
		}
	}
}

//给指定分组发送消息
func (h *Hub) SendGroupMsg(groupClientList []*Client, msg []byte) {
	for _, client := range groupClientList {
		client.SendMsg(msg)
	}
}
