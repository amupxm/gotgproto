// GoTGProto Generic Helpers
// WARNING: This file is autogenerated, please DO NOT EDIT

package generic

import (
	"github.com/amupxm/gotgproto/ext"
	"github.com/amupxm/gotgproto/types"
	"github.com/gotd/td/tg"
)

type ChatUnion interface {
	int | int64 | string
}

func getIdByUnion[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion) (int64, error) {
	switch val := any(chat).(type) {
	case string:
		username := val
		peer := ctx.PeerStorage.GetPeerByUsername(username)
		if peer.ID != 0 {
			return peer.ID, nil
		}
		chat, err := ctx.ResolveUsername(username)
		return chat.GetID(), err
	case int64:
		return val, nil
	case int:
		return int64(val), nil
	}
	// Unreachable
	return 0, nil
}

// SendMessage is a generic helper for ext.Context.SendMessage method.
func SendMessage[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesSendMessageRequest) (*types.Message, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.SendMessage(chatId, request)
}

// SendMedia is a generic helper for ext.Context.SendMedia method.
func SendMedia[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesSendMediaRequest) (*types.Message, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.SendMedia(chatId, request)
}

// GetInlineBotResults is a generic helper for ext.Context.GetInlineBotResults method.
func GetInlineBotResults[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, botUsername string, request *tg.MessagesGetInlineBotResultsRequest) (*tg.MessagesBotResults, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.GetInlineBotResults(chatId, botUsername, request)
}

// SendInlineBotResult is a generic helper for ext.Context.SendInlineBotResult method.
func SendInlineBotResult[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesSendInlineBotResultRequest) (tg.UpdatesClass, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.SendInlineBotResult(chatId, request)
}

// SendReaction is a generic helper for ext.Context.SendReaction method.
func SendReaction[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesSendReactionRequest) (*types.Message, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.SendReaction(chatId, request)
}

// SendMultiMedia is a generic helper for ext.Context.SendMultiMedia method.
func SendMultiMedia[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesSendMultiMediaRequest) (*types.Message, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.SendMultiMedia(chatId, request)
}

// EditMessage is a generic helper for ext.Context.EditMessage method.
func EditMessage[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, request *tg.MessagesEditMessageRequest) (*types.Message, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.EditMessage(chatId, request)
}

// GetChat is a generic helper for ext.Context.GetChat method.
func GetChat[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion) (tg.ChatFullClass, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.GetChat(chatId)
}

// GetUser is a generic helper for ext.Context.GetUser method.
func GetUser[chatUnion ChatUnion](ctx *ext.Context, user chatUnion) (*tg.UserFull, error) {

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return nil, err
	}

	return ctx.GetUser(userId)
}

// GetMessages is a generic helper for ext.Context.GetMessages method.
func GetMessages[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, messageIds []tg.InputMessageClass) ([]tg.MessageClass, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	return ctx.GetMessages(chatId, messageIds)
}

// BanChatMember is a generic helper for ext.Context.BanChatMember method.
func BanChatMember[chatUnion ChatUnion](ctx *ext.Context, chat, user chatUnion, untilDate int) (tg.UpdatesClass, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return nil, err
	}

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return nil, err
	}

	return ctx.BanChatMember(chatId, userId, untilDate)
}

// UnbanChatMember is a generic helper for ext.Context.UnbanChatMember method.
func UnbanChatMember[chatUnion ChatUnion](ctx *ext.Context, chat, user chatUnion) (bool, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return false, err
	}

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return false, err
	}

	return ctx.UnbanChatMember(chatId, userId)
}

// AddChatMembers is a generic helper for ext.Context.AddChatMembers method.
func AddChatMembers[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, userIds []int64, forwardLimit int) (bool, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return false, err
	}

	return ctx.AddChatMembers(chatId, userIds, forwardLimit)
}

// DeleteMessages is a generic helper for ext.Context.DeleteMessages method.
func DeleteMessages[chatUnion ChatUnion](ctx *ext.Context, chat chatUnion, messageIDs []int) error {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return err
	}

	return ctx.DeleteMessages(chatId, messageIDs)
}

// PromoteChatMember is a generic helper for ext.Context.PromoteChatMember method.
func PromoteChatMember[chatUnion ChatUnion](ctx *ext.Context, chat, user chatUnion, opts *ext.EditAdminOpts) (bool, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return false, err
	}

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return false, err
	}

	return ctx.PromoteChatMember(chatId, userId, opts)
}

// DemoteChatMember is a generic helper for ext.Context.DemoteChatMember method.
func DemoteChatMember[chatUnion ChatUnion](ctx *ext.Context, chat, user chatUnion, opts *ext.EditAdminOpts) (bool, error) {

	chatId, err := getIdByUnion(ctx, chat)
	if err != nil {
		return false, err
	}

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return false, err
	}

	return ctx.DemoteChatMember(chatId, userId, opts)
}

// GetUserProfilePhotos is a generic helper for ext.Context.GetUserProfilePhotos method.
func GetUserProfilePhotos[chatUnion ChatUnion](ctx *ext.Context, user chatUnion, opts *tg.PhotosGetUserPhotosRequest) ([]tg.PhotoClass, error) {

	userId, err := getIdByUnion(ctx, user)
	if err != nil {
		return nil, err
	}

	return ctx.GetUserProfilePhotos(userId, opts)
}
