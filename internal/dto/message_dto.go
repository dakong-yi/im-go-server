package dto

type V2TimMessageListResult struct {
	IsFinished  bool            `json:"isFinished"`
	MessageList []*V2TimMessage `json:"messageList"`
}

type V2TimMessage struct {
	MsgID                     string              `json:"msgID"`
	Timestamp                 int64               `json:"timestamp"`
	Progress                  int                 `json:"progress"`
	Sender                    string              `json:"sender"`
	NickName                  string              `json:"nickName"`
	FriendRemark              string              `json:"friendRemark"`
	FaceUrl                   string              `json:"faceUrl"`
	NameCard                  string              `json:"nameCard"`
	GroupID                   string              `json:"groupID"`
	UserID                    string              `json:"userID"`
	Status                    int                 `json:"status"`
	ElemType                  int                 `json:"elemType"`
	TextElem                  *V2TimTextElem      `json:"textElem,omitempty"`
	CustomElem                *V2TimCustomElem    `json:"customElem,omitempty"`
	ImageElem                 *V2TimImageElem     `json:"imageElem,omitempty"`
	SoundElem                 *V2TimSoundElem     `json:"soundElem,omitempty"`
	VideoElem                 *V2TimVideoElem     `json:"videoElem,omitempty"`
	FileElem                  *V2TimFileElem      `json:"fileElem,omitempty"`
	LocationElem              *V2TimLocationElem  `json:"locationElem,omitempty"`
	FaceElem                  *V2TimFaceElem      `json:"faceElem,omitempty"`
	GroupTipsElem             *V2TimGroupTipsElem `json:"groupTipsElem,omitempty"`
	MergerElem                *V2TimMergerElem    `json:"mergerElem,omitempty"`
	LocalCustomData           string              `json:"localCustomData"`
	LocalCustomInt            int                 `json:"localCustomInt"`
	CloudCustomData           string              `json:"cloudCustomData"`
	IsSelf                    bool                `json:"isSelf"`
	IsRead                    bool                `json:"isRead"`
	IsPeerRead                bool                `json:"isPeerRead"`
	Priority                  int                 `json:"priority"`
	OfflinePushInfo           *OfflinePushInfo    `json:"offlinePushInfo,omitempty"`
	GroupAtUserList           []string            `json:"groupAtUserList,omitempty"`
	Seq                       string              `json:"seq"`
	Random                    int                 `json:"random"`
	IsExcludedFromUnreadCount bool                `json:"isExcludedFromUnreadCount"`
	IsExcludedFromLastMessage bool                `json:"isExcludedFromLastMessage"`
	IsSupportMessageExtension bool                `json:"isSupportMessageExtension"`
	MessageFromWeb            string              `json:"messageFromWeb"`
	ID                        string              `json:"id"`
	NeedReadReceipt           bool                `json:"needReadReceipt"`
}

type V2TimTextElem struct {
	Text string `json:"text"`
}

type V2TimCustomElem struct {
	Data      string `json:"data"`
	Desc      string `json:"desc"`
	Extension string `json:"extension"`
}

type V2TimImageElem struct {
	Path      string       `json:"path"`
	ImageList []V2TimImage `json:"imageList"`
}

type V2TimImage struct {
	UUID     *string `json:"uuid"`
	Type     *int    `json:"type"`
	Size     *int    `json:"size"`
	Width    *int    `json:"width"`
	Height   *int    `json:"height"`
	URL      *string `json:"url"`
	LocalURL *string `json:"localUrl"`
}

type V2TimSoundElem struct {
	Path     string `json:"path"`
	UUID     string `json:"UUID"`
	DataSize int    `json:"dataSize"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LocalURL string `json:"localUrl"`
}

type V2TimVideoElem struct {
	VideoPath        *string `json:"videoPath"`
	UUID             *string `json:"UUID"`
	VideoSize        *int    `json:"videoSize"`
	Duration         *int    `json:"duration"`
	SnapshotPath     *string `json:"snapshotPath"`
	SnapshotUUID     *string `json:"snapshotUUID"`
	SnapshotSize     *int    `json:"snapshotSize"`
	SnapshotWidth    *int    `json:"snapshotWidth"`
	SnapshotHeight   *int    `json:"snapshotHeight"`
	VideoUrl         *string `json:"videoUrl"`
	SnapshotUrl      *string `json:"snapshotUrl"`
	LocalVideoUrl    *string `json:"localVideoUrl"`
	LocalSnapshotUrl *string `json:"localSnapshotUrl"`
}

type V2TimFileElem struct {
	Path     *string `json:"path,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	UUID     *string `json:"UUID,omitempty"`
	URL      *string `json:"url,omitempty"`
	FileSize *int    `json:"fileSize,omitempty"`
	LocalURL *string `json:"localUrl,omitempty"`
}

type V2TimLocationElem struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Desc      *string  `json:"desc"`
}

type V2TimFaceElem struct {
	Index int    `json:"index"`
	Data  string `json:"data"`
}

type V2TimGroupTipsElem struct {
}

type V2TimMergerElem struct {
	IsLayersOverLimit bool     `json:"isLayersOverLimit"`
	Title             string   `json:"title"`
	AbstractList      []string `json:"abstractList"`
}

type OfflinePushInfo struct {
	Title                     *string `json:"title"`
	Desc                      *string `json:"desc"`
	Ext                       *string `json:"ext"`
	DisablePush               *bool   `json:"disablePush"`
	IOSSound                  *string `json:"iOSSound"`
	IgnoreIOSBadge            *bool   `json:"ignoreIOSBadge"`
	AndroidOPPOChannelID      *string `json:"androidOPPOChannelID"`
	AndroidVIVOClassification *int    `json:"androidVIVOClassification"`
	AndroidSound              *string `json:"androidSound"`
	AndroidFCMChannelID       *string `json:"androidFCMChannelID"`
	AndroidXiaoMiChannelID    *string `json:"androidXiaoMiChannelID"`
	IOSPushType               *int    `json:"iOSPushType"`
	AndroidHuaWeiCategory     *string `json:"androidHuaWeiCategory"`
}

type V2TimMsgCreateInfoResult struct {
	ID          string        `json:"id"`
	MessageInfo *V2TimMessage `json:"messageInfo"`
}
