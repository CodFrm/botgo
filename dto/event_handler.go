package dto

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild           GuildEventHandler
	GuildMember     GuildMemberEventHandler
	Channel         ChannelEventHandler
	Message         MessageEventHandler
	MessageReaction MessageReactionEventHandler
	ATMessage       ATMessageEventHandler
	DirectMessage   DirectMessageEventHandler
	Audio           AudioEventHandler
	MessageAudit    MessageAuditEventHandler
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *WSPayload, data *WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *WSPayload, data *WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *WSPayload, data *WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *WSPayload, data *WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *WSPayload, data *WSMessageData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *WSPayload, data *WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *WSPayload, data *WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *WSPayload, data *WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *WSPayload, data *WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *WSPayload, data *WSMessageAuditData) error
