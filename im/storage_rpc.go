/**
 * Copyright (c) 2014-2015, GoBelieve
 * All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
 */

package main

type PeerMessage struct {
	AppID    int64
	Uid      int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type PeerGroupMessage struct {
	AppID    int64
	Members  []int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type GroupMessage struct {
	AppID    int64
	GroupID  int64
	DeviceID int64
	Cmd      int32
	Raw      []byte
}

type HistoryMessage struct {
	MsgID    int64
	DeviceID int64 //消息发送者所在的设备ID
	Cmd      int32
	Raw      []byte
}

type PeerHistoryMessage struct {
	Messages  []*HistoryMessage
	LastMsgID int64
	HasMore   bool
}

type GroupHistoryMessage PeerHistoryMessage

type SyncHistory struct {
	AppID     int64
	Uid       int64
	DeviceID  int64
	LastMsgID int64
}

type SyncGroupHistory struct {
	AppID     int64
	Uid       int64
	DeviceID  int64
	GroupID   int64
	LastMsgID int64
	Timestamp int32
}

type HistoryRequest struct {
	AppID int64
	Uid   int64
	LastMsgID int64
	Limit int32
}

func SyncMessageInterface(addr string, sync_key *SyncHistory) *PeerHistoryMessage {
	return nil
}

func SyncGroupMessageInterface(addr string, sync_key *SyncGroupHistory) *GroupHistoryMessage {
	return nil
}

func SavePeerMessageInterface(addr string, m *PeerMessage) ([2]int64, error) {
	return [2]int64{}, nil
}

func SavePeerGroupMessageInterface(addr string, m *PeerGroupMessage) ([]int64, error) {
	return nil, nil
}

func SaveGroupMessageInterface(addr string, m *GroupMessage) ([2]int64, error) {
	return [2]int64{}, nil
}

//获取是否接收到新消息,只会返回0/1
func GetNewCountInterface(addr string, s *SyncHistory) (int64, error) {
	return 0, nil
}

func GetLatestMessageInterface(addr string, r *HistoryRequest) []*HistoryMessage {
	return nil
}
