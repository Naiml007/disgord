package disgord

// Code generated - This file has been automatically generated by generate/events/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt reactor.go and event/events.go and run `go generate`

//////////////////////////////////////////////////////
//
// Helper funcs
//
//////////////////////////////////////////////////////

func defineResource(evt string) (resource evtResource) {
	switch evt {

	case EvtChannelCreate:
		resource = &ChannelCreate{}
	case EvtChannelDelete:
		resource = &ChannelDelete{}
	case EvtChannelPinsUpdate:
		resource = &ChannelPinsUpdate{}
	case EvtChannelUpdate:
		resource = &ChannelUpdate{}
	case EvtGuildBanAdd:
		resource = &GuildBanAdd{}
	case EvtGuildBanRemove:
		resource = &GuildBanRemove{}
	case EvtGuildCreate:
		resource = &GuildCreate{}
	case EvtGuildDelete:
		resource = &GuildDelete{}
	case EvtGuildEmojisUpdate:
		resource = &GuildEmojisUpdate{}
	case EvtGuildIntegrationsUpdate:
		resource = &GuildIntegrationsUpdate{}
	case EvtGuildMemberAdd:
		resource = &GuildMemberAdd{}
	case EvtGuildMemberRemove:
		resource = &GuildMemberRemove{}
	case EvtGuildMemberUpdate:
		resource = &GuildMemberUpdate{}
	case EvtGuildMembersChunk:
		resource = &GuildMembersChunk{}
	case EvtGuildRoleCreate:
		resource = &GuildRoleCreate{}
	case EvtGuildRoleDelete:
		resource = &GuildRoleDelete{}
	case EvtGuildRoleUpdate:
		resource = &GuildRoleUpdate{}
	case EvtGuildUpdate:
		resource = &GuildUpdate{}
	case EvtInviteCreate:
		resource = &InviteCreate{}
	case EvtInviteDelete:
		resource = &InviteDelete{}
	case EvtMessageCreate:
		resource = &MessageCreate{}
	case EvtMessageDelete:
		resource = &MessageDelete{}
	case EvtMessageDeleteBulk:
		resource = &MessageDeleteBulk{}
	case EvtMessageReactionAdd:
		resource = &MessageReactionAdd{}
	case EvtMessageReactionRemove:
		resource = &MessageReactionRemove{}
	case EvtMessageReactionRemoveAll:
		resource = &MessageReactionRemoveAll{}
	case EvtMessageReactionRemoveEmoji:
		resource = &MessageReactionRemoveEmoji{}
	case EvtMessageUpdate:
		resource = &MessageUpdate{}
	case EvtPresenceUpdate:
		resource = &PresenceUpdate{}
	case EvtReady:
		resource = &Ready{}
	case EvtResumed:
		resource = &Resumed{}
	case EvtTypingStart:
		resource = &TypingStart{}
	case EvtUserUpdate:
		resource = &UserUpdate{}
	case EvtVoiceServerUpdate:
		resource = &VoiceServerUpdate{}
	case EvtVoiceStateUpdate:
		resource = &VoiceStateUpdate{}
	case EvtWebhooksUpdate:
		resource = &WebhooksUpdate{}
	}

	return resource
}

func isHandler(h Handler) (ok bool) {
	switch h.(type) {
	case HandlerSimple:
		ok = true
	case HandlerSimplest:
		ok = true
	case chan interface{}:
		ok = true
	case HandlerChannelCreate:
		ok = true
	case chan *ChannelCreate:
		ok = true
	case HandlerChannelDelete:
		ok = true
	case chan *ChannelDelete:
		ok = true
	case HandlerChannelPinsUpdate:
		ok = true
	case chan *ChannelPinsUpdate:
		ok = true
	case HandlerChannelUpdate:
		ok = true
	case chan *ChannelUpdate:
		ok = true
	case HandlerGuildBanAdd:
		ok = true
	case chan *GuildBanAdd:
		ok = true
	case HandlerGuildBanRemove:
		ok = true
	case chan *GuildBanRemove:
		ok = true
	case HandlerGuildCreate:
		ok = true
	case chan *GuildCreate:
		ok = true
	case HandlerGuildDelete:
		ok = true
	case chan *GuildDelete:
		ok = true
	case HandlerGuildEmojisUpdate:
		ok = true
	case chan *GuildEmojisUpdate:
		ok = true
	case HandlerGuildIntegrationsUpdate:
		ok = true
	case chan *GuildIntegrationsUpdate:
		ok = true
	case HandlerGuildMemberAdd:
		ok = true
	case chan *GuildMemberAdd:
		ok = true
	case HandlerGuildMemberRemove:
		ok = true
	case chan *GuildMemberRemove:
		ok = true
	case HandlerGuildMemberUpdate:
		ok = true
	case chan *GuildMemberUpdate:
		ok = true
	case HandlerGuildMembersChunk:
		ok = true
	case chan *GuildMembersChunk:
		ok = true
	case HandlerGuildRoleCreate:
		ok = true
	case chan *GuildRoleCreate:
		ok = true
	case HandlerGuildRoleDelete:
		ok = true
	case chan *GuildRoleDelete:
		ok = true
	case HandlerGuildRoleUpdate:
		ok = true
	case chan *GuildRoleUpdate:
		ok = true
	case HandlerGuildUpdate:
		ok = true
	case chan *GuildUpdate:
		ok = true
	case HandlerInviteCreate:
		ok = true
	case chan *InviteCreate:
		ok = true
	case HandlerInviteDelete:
		ok = true
	case chan *InviteDelete:
		ok = true
	case HandlerMessageCreate:
		ok = true
	case chan *MessageCreate:
		ok = true
	case HandlerMessageDelete:
		ok = true
	case chan *MessageDelete:
		ok = true
	case HandlerMessageDeleteBulk:
		ok = true
	case chan *MessageDeleteBulk:
		ok = true
	case HandlerMessageReactionAdd:
		ok = true
	case chan *MessageReactionAdd:
		ok = true
	case HandlerMessageReactionRemove:
		ok = true
	case chan *MessageReactionRemove:
		ok = true
	case HandlerMessageReactionRemoveAll:
		ok = true
	case chan *MessageReactionRemoveAll:
		ok = true
	case HandlerMessageReactionRemoveEmoji:
		ok = true
	case chan *MessageReactionRemoveEmoji:
		ok = true
	case HandlerMessageUpdate:
		ok = true
	case chan *MessageUpdate:
		ok = true
	case HandlerPresenceUpdate:
		ok = true
	case chan *PresenceUpdate:
		ok = true
	case HandlerReady:
		ok = true
	case chan *Ready:
		ok = true
	case HandlerResumed:
		ok = true
	case chan *Resumed:
		ok = true
	case HandlerTypingStart:
		ok = true
	case chan *TypingStart:
		ok = true
	case HandlerUserUpdate:
		ok = true
	case chan *UserUpdate:
		ok = true
	case HandlerVoiceServerUpdate:
		ok = true
	case chan *VoiceServerUpdate:
		ok = true
	case HandlerVoiceStateUpdate:
		ok = true
	case chan *VoiceStateUpdate:
		ok = true
	case HandlerWebhooksUpdate:
		ok = true
	case chan *WebhooksUpdate:
		ok = true
	}
	return ok
}

func closeChannel(channel interface{}) {
	switch t := channel.(type) {
	case chan interface{}:
		close(t)
	case chan *ChannelCreate:
		close(t)
	case chan *ChannelDelete:
		close(t)
	case chan *ChannelPinsUpdate:
		close(t)
	case chan *ChannelUpdate:
		close(t)
	case chan *GuildBanAdd:
		close(t)
	case chan *GuildBanRemove:
		close(t)
	case chan *GuildCreate:
		close(t)
	case chan *GuildDelete:
		close(t)
	case chan *GuildEmojisUpdate:
		close(t)
	case chan *GuildIntegrationsUpdate:
		close(t)
	case chan *GuildMemberAdd:
		close(t)
	case chan *GuildMemberRemove:
		close(t)
	case chan *GuildMemberUpdate:
		close(t)
	case chan *GuildMembersChunk:
		close(t)
	case chan *GuildRoleCreate:
		close(t)
	case chan *GuildRoleDelete:
		close(t)
	case chan *GuildRoleUpdate:
		close(t)
	case chan *GuildUpdate:
		close(t)
	case chan *InviteCreate:
		close(t)
	case chan *InviteDelete:
		close(t)
	case chan *MessageCreate:
		close(t)
	case chan *MessageDelete:
		close(t)
	case chan *MessageDeleteBulk:
		close(t)
	case chan *MessageReactionAdd:
		close(t)
	case chan *MessageReactionRemove:
		close(t)
	case chan *MessageReactionRemoveAll:
		close(t)
	case chan *MessageReactionRemoveEmoji:
		close(t)
	case chan *MessageUpdate:
		close(t)
	case chan *PresenceUpdate:
		close(t)
	case chan *Ready:
		close(t)
	case chan *Resumed:
		close(t)
	case chan *TypingStart:
		close(t)
	case chan *UserUpdate:
		close(t)
	case chan *VoiceServerUpdate:
		close(t)
	case chan *VoiceStateUpdate:
		close(t)
	case chan *WebhooksUpdate:
		close(t)
	}
}

//////////////////////////////////////////////////////
//
// Dispatcher: contructor + repetitive methods
//
//////////////////////////////////////////////////////

// newDispatcher construct a Dispatch object for reacting to web socket events
// from discord
func newDispatcher() *dispatcher {
	d := &dispatcher{
		handlerSpecs: make(map[string][]*handlerSpec),
		shutdown:     make(chan struct{}),
	}

	return d
}

func (d *dispatcher) trigger(h Handler, evt resource) {
	switch t := h.(type) {
	case HandlerSimple:
		t(d.session)
	case HandlerSimplest:
		t()
	case chan interface{}:
		t <- evt
	case chan<- interface{}:
		t <- evt
	case HandlerChannelCreate:
		t(d.session, evt.(*ChannelCreate))
	case chan *ChannelCreate:
		t <- evt.(*ChannelCreate)
	case chan<- *ChannelCreate:
		t <- evt.(*ChannelCreate)
	case HandlerChannelDelete:
		t(d.session, evt.(*ChannelDelete))
	case chan *ChannelDelete:
		t <- evt.(*ChannelDelete)
	case chan<- *ChannelDelete:
		t <- evt.(*ChannelDelete)
	case HandlerChannelPinsUpdate:
		t(d.session, evt.(*ChannelPinsUpdate))
	case chan *ChannelPinsUpdate:
		t <- evt.(*ChannelPinsUpdate)
	case chan<- *ChannelPinsUpdate:
		t <- evt.(*ChannelPinsUpdate)
	case HandlerChannelUpdate:
		t(d.session, evt.(*ChannelUpdate))
	case chan *ChannelUpdate:
		t <- evt.(*ChannelUpdate)
	case chan<- *ChannelUpdate:
		t <- evt.(*ChannelUpdate)
	case HandlerGuildBanAdd:
		t(d.session, evt.(*GuildBanAdd))
	case chan *GuildBanAdd:
		t <- evt.(*GuildBanAdd)
	case chan<- *GuildBanAdd:
		t <- evt.(*GuildBanAdd)
	case HandlerGuildBanRemove:
		t(d.session, evt.(*GuildBanRemove))
	case chan *GuildBanRemove:
		t <- evt.(*GuildBanRemove)
	case chan<- *GuildBanRemove:
		t <- evt.(*GuildBanRemove)
	case HandlerGuildCreate:
		t(d.session, evt.(*GuildCreate))
	case chan *GuildCreate:
		t <- evt.(*GuildCreate)
	case chan<- *GuildCreate:
		t <- evt.(*GuildCreate)
	case HandlerGuildDelete:
		t(d.session, evt.(*GuildDelete))
	case chan *GuildDelete:
		t <- evt.(*GuildDelete)
	case chan<- *GuildDelete:
		t <- evt.(*GuildDelete)
	case HandlerGuildEmojisUpdate:
		t(d.session, evt.(*GuildEmojisUpdate))
	case chan *GuildEmojisUpdate:
		t <- evt.(*GuildEmojisUpdate)
	case chan<- *GuildEmojisUpdate:
		t <- evt.(*GuildEmojisUpdate)
	case HandlerGuildIntegrationsUpdate:
		t(d.session, evt.(*GuildIntegrationsUpdate))
	case chan *GuildIntegrationsUpdate:
		t <- evt.(*GuildIntegrationsUpdate)
	case chan<- *GuildIntegrationsUpdate:
		t <- evt.(*GuildIntegrationsUpdate)
	case HandlerGuildMemberAdd:
		t(d.session, evt.(*GuildMemberAdd))
	case chan *GuildMemberAdd:
		t <- evt.(*GuildMemberAdd)
	case chan<- *GuildMemberAdd:
		t <- evt.(*GuildMemberAdd)
	case HandlerGuildMemberRemove:
		t(d.session, evt.(*GuildMemberRemove))
	case chan *GuildMemberRemove:
		t <- evt.(*GuildMemberRemove)
	case chan<- *GuildMemberRemove:
		t <- evt.(*GuildMemberRemove)
	case HandlerGuildMemberUpdate:
		t(d.session, evt.(*GuildMemberUpdate))
	case chan *GuildMemberUpdate:
		t <- evt.(*GuildMemberUpdate)
	case chan<- *GuildMemberUpdate:
		t <- evt.(*GuildMemberUpdate)
	case HandlerGuildMembersChunk:
		t(d.session, evt.(*GuildMembersChunk))
	case chan *GuildMembersChunk:
		t <- evt.(*GuildMembersChunk)
	case chan<- *GuildMembersChunk:
		t <- evt.(*GuildMembersChunk)
	case HandlerGuildRoleCreate:
		t(d.session, evt.(*GuildRoleCreate))
	case chan *GuildRoleCreate:
		t <- evt.(*GuildRoleCreate)
	case chan<- *GuildRoleCreate:
		t <- evt.(*GuildRoleCreate)
	case HandlerGuildRoleDelete:
		t(d.session, evt.(*GuildRoleDelete))
	case chan *GuildRoleDelete:
		t <- evt.(*GuildRoleDelete)
	case chan<- *GuildRoleDelete:
		t <- evt.(*GuildRoleDelete)
	case HandlerGuildRoleUpdate:
		t(d.session, evt.(*GuildRoleUpdate))
	case chan *GuildRoleUpdate:
		t <- evt.(*GuildRoleUpdate)
	case chan<- *GuildRoleUpdate:
		t <- evt.(*GuildRoleUpdate)
	case HandlerGuildUpdate:
		t(d.session, evt.(*GuildUpdate))
	case chan *GuildUpdate:
		t <- evt.(*GuildUpdate)
	case chan<- *GuildUpdate:
		t <- evt.(*GuildUpdate)
	case HandlerInviteCreate:
		t(d.session, evt.(*InviteCreate))
	case chan *InviteCreate:
		t <- evt.(*InviteCreate)
	case chan<- *InviteCreate:
		t <- evt.(*InviteCreate)
	case HandlerInviteDelete:
		t(d.session, evt.(*InviteDelete))
	case chan *InviteDelete:
		t <- evt.(*InviteDelete)
	case chan<- *InviteDelete:
		t <- evt.(*InviteDelete)
	case HandlerMessageCreate:
		t(d.session, evt.(*MessageCreate))
	case chan *MessageCreate:
		t <- evt.(*MessageCreate)
	case chan<- *MessageCreate:
		t <- evt.(*MessageCreate)
	case HandlerMessageDelete:
		t(d.session, evt.(*MessageDelete))
	case chan *MessageDelete:
		t <- evt.(*MessageDelete)
	case chan<- *MessageDelete:
		t <- evt.(*MessageDelete)
	case HandlerMessageDeleteBulk:
		t(d.session, evt.(*MessageDeleteBulk))
	case chan *MessageDeleteBulk:
		t <- evt.(*MessageDeleteBulk)
	case chan<- *MessageDeleteBulk:
		t <- evt.(*MessageDeleteBulk)
	case HandlerMessageReactionAdd:
		t(d.session, evt.(*MessageReactionAdd))
	case chan *MessageReactionAdd:
		t <- evt.(*MessageReactionAdd)
	case chan<- *MessageReactionAdd:
		t <- evt.(*MessageReactionAdd)
	case HandlerMessageReactionRemove:
		t(d.session, evt.(*MessageReactionRemove))
	case chan *MessageReactionRemove:
		t <- evt.(*MessageReactionRemove)
	case chan<- *MessageReactionRemove:
		t <- evt.(*MessageReactionRemove)
	case HandlerMessageReactionRemoveAll:
		t(d.session, evt.(*MessageReactionRemoveAll))
	case chan *MessageReactionRemoveAll:
		t <- evt.(*MessageReactionRemoveAll)
	case chan<- *MessageReactionRemoveAll:
		t <- evt.(*MessageReactionRemoveAll)
	case HandlerMessageReactionRemoveEmoji:
		t(d.session, evt.(*MessageReactionRemoveEmoji))
	case chan *MessageReactionRemoveEmoji:
		t <- evt.(*MessageReactionRemoveEmoji)
	case chan<- *MessageReactionRemoveEmoji:
		t <- evt.(*MessageReactionRemoveEmoji)
	case HandlerMessageUpdate:
		t(d.session, evt.(*MessageUpdate))
	case chan *MessageUpdate:
		t <- evt.(*MessageUpdate)
	case chan<- *MessageUpdate:
		t <- evt.(*MessageUpdate)
	case HandlerPresenceUpdate:
		t(d.session, evt.(*PresenceUpdate))
	case chan *PresenceUpdate:
		t <- evt.(*PresenceUpdate)
	case chan<- *PresenceUpdate:
		t <- evt.(*PresenceUpdate)
	case HandlerReady:
		t(d.session, evt.(*Ready))
	case chan *Ready:
		t <- evt.(*Ready)
	case chan<- *Ready:
		t <- evt.(*Ready)
	case HandlerResumed:
		t(d.session, evt.(*Resumed))
	case chan *Resumed:
		t <- evt.(*Resumed)
	case chan<- *Resumed:
		t <- evt.(*Resumed)
	case HandlerTypingStart:
		t(d.session, evt.(*TypingStart))
	case chan *TypingStart:
		t <- evt.(*TypingStart)
	case chan<- *TypingStart:
		t <- evt.(*TypingStart)
	case HandlerUserUpdate:
		t(d.session, evt.(*UserUpdate))
	case chan *UserUpdate:
		t <- evt.(*UserUpdate)
	case chan<- *UserUpdate:
		t <- evt.(*UserUpdate)
	case HandlerVoiceServerUpdate:
		t(d.session, evt.(*VoiceServerUpdate))
	case chan *VoiceServerUpdate:
		t <- evt.(*VoiceServerUpdate)
	case chan<- *VoiceServerUpdate:
		t <- evt.(*VoiceServerUpdate)
	case HandlerVoiceStateUpdate:
		t(d.session, evt.(*VoiceStateUpdate))
	case chan *VoiceStateUpdate:
		t <- evt.(*VoiceStateUpdate)
	case chan<- *VoiceStateUpdate:
		t <- evt.(*VoiceStateUpdate)
	case HandlerWebhooksUpdate:
		t(d.session, evt.(*WebhooksUpdate))
	case chan *WebhooksUpdate:
		t <- evt.(*WebhooksUpdate)
	case chan<- *WebhooksUpdate:
		t <- evt.(*WebhooksUpdate)
	}
}

//////////////////////////////////////////////////////
//
// Handler Signatures
//
//////////////////////////////////////////////////////

// these "simple" handler can be used, if you don't care about the actual event data
type HandlerSimplest = func()
type HandlerSimple = func(Session)

// HandlerChannelCreate is triggered by ChannelCreate events
type HandlerChannelCreate = func(s Session, h *ChannelCreate)

// HandlerChannelDelete is triggered by ChannelDelete events
type HandlerChannelDelete = func(s Session, h *ChannelDelete)

// HandlerChannelPinsUpdate is triggered by ChannelPinsUpdate events
type HandlerChannelPinsUpdate = func(s Session, h *ChannelPinsUpdate)

// HandlerChannelUpdate is triggered by ChannelUpdate events
type HandlerChannelUpdate = func(s Session, h *ChannelUpdate)

// HandlerGuildBanAdd is triggered by GuildBanAdd events
type HandlerGuildBanAdd = func(s Session, h *GuildBanAdd)

// HandlerGuildBanRemove is triggered by GuildBanRemove events
type HandlerGuildBanRemove = func(s Session, h *GuildBanRemove)

// HandlerGuildCreate is triggered by GuildCreate events
type HandlerGuildCreate = func(s Session, h *GuildCreate)

// HandlerGuildDelete is triggered by GuildDelete events
type HandlerGuildDelete = func(s Session, h *GuildDelete)

// HandlerGuildEmojisUpdate is triggered by GuildEmojisUpdate events
type HandlerGuildEmojisUpdate = func(s Session, h *GuildEmojisUpdate)

// HandlerGuildIntegrationsUpdate is triggered by GuildIntegrationsUpdate events
type HandlerGuildIntegrationsUpdate = func(s Session, h *GuildIntegrationsUpdate)

// HandlerGuildMemberAdd is triggered by GuildMemberAdd events
type HandlerGuildMemberAdd = func(s Session, h *GuildMemberAdd)

// HandlerGuildMemberRemove is triggered by GuildMemberRemove events
type HandlerGuildMemberRemove = func(s Session, h *GuildMemberRemove)

// HandlerGuildMemberUpdate is triggered by GuildMemberUpdate events
type HandlerGuildMemberUpdate = func(s Session, h *GuildMemberUpdate)

// HandlerGuildMembersChunk is triggered by GuildMembersChunk events
type HandlerGuildMembersChunk = func(s Session, h *GuildMembersChunk)

// HandlerGuildRoleCreate is triggered by GuildRoleCreate events
type HandlerGuildRoleCreate = func(s Session, h *GuildRoleCreate)

// HandlerGuildRoleDelete is triggered by GuildRoleDelete events
type HandlerGuildRoleDelete = func(s Session, h *GuildRoleDelete)

// HandlerGuildRoleUpdate is triggered by GuildRoleUpdate events
type HandlerGuildRoleUpdate = func(s Session, h *GuildRoleUpdate)

// HandlerGuildUpdate is triggered by GuildUpdate events
type HandlerGuildUpdate = func(s Session, h *GuildUpdate)

// HandlerInviteCreate is triggered by InviteCreate events
type HandlerInviteCreate = func(s Session, h *InviteCreate)

// HandlerInviteDelete is triggered by InviteDelete events
type HandlerInviteDelete = func(s Session, h *InviteDelete)

// HandlerMessageCreate is triggered by MessageCreate events
type HandlerMessageCreate = func(s Session, h *MessageCreate)

// HandlerMessageDelete is triggered by MessageDelete events
type HandlerMessageDelete = func(s Session, h *MessageDelete)

// HandlerMessageDeleteBulk is triggered by MessageDeleteBulk events
type HandlerMessageDeleteBulk = func(s Session, h *MessageDeleteBulk)

// HandlerMessageReactionAdd is triggered by MessageReactionAdd events
type HandlerMessageReactionAdd = func(s Session, h *MessageReactionAdd)

// HandlerMessageReactionRemove is triggered by MessageReactionRemove events
type HandlerMessageReactionRemove = func(s Session, h *MessageReactionRemove)

// HandlerMessageReactionRemoveAll is triggered by MessageReactionRemoveAll events
type HandlerMessageReactionRemoveAll = func(s Session, h *MessageReactionRemoveAll)

// HandlerMessageReactionRemoveEmoji is triggered by MessageReactionRemoveEmoji events
type HandlerMessageReactionRemoveEmoji = func(s Session, h *MessageReactionRemoveEmoji)

// HandlerMessageUpdate is triggered by MessageUpdate events
type HandlerMessageUpdate = func(s Session, h *MessageUpdate)

// HandlerPresenceUpdate is triggered by PresenceUpdate events
type HandlerPresenceUpdate = func(s Session, h *PresenceUpdate)

// HandlerReady is triggered by Ready events
type HandlerReady = func(s Session, h *Ready)

// HandlerResumed is triggered by Resumed events
type HandlerResumed = func(s Session, h *Resumed)

// HandlerTypingStart is triggered by TypingStart events
type HandlerTypingStart = func(s Session, h *TypingStart)

// HandlerUserUpdate is triggered by UserUpdate events
type HandlerUserUpdate = func(s Session, h *UserUpdate)

// HandlerVoiceServerUpdate is triggered by VoiceServerUpdate events
type HandlerVoiceServerUpdate = func(s Session, h *VoiceServerUpdate)

// HandlerVoiceStateUpdate is triggered by VoiceStateUpdate events
type HandlerVoiceStateUpdate = func(s Session, h *VoiceStateUpdate)

// HandlerWebhooksUpdate is triggered by WebhooksUpdate events
type HandlerWebhooksUpdate = func(s Session, h *WebhooksUpdate)
