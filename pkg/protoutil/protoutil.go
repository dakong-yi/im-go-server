package protoutil

// import (
// 	"encoding/binary"
// 	"io"

// 	"google.golang.org/protobuf/proto"
// 	// "google.golang.org/protobuf/proto"
// )

// // TODO proto格式传输暂未实现
// func WriteProtoMessage(conn io.Writer, message proto.Message) error {
// 	data, err := proto.Marshal(message)
// 	if err != nil {
// 		return err
// 	}
// 	length := make([]byte, 4)
// 	binary.BigEndian.PutUint32(length, uint32(len(data)))
// 	_, err = conn.Write(length)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = conn.Write(data)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func ReadProtoMessage(conn io.Reader, message proto.Message) error {
// 	lengthBytes := make([]byte, 4)
// 	_, err := io.ReadFull(conn, lengthBytes)
// 	if err != nil {
// 		return err
// 	}
// 	length := binary.BigEndian.Uint32(lengthBytes)
// 	data := make([]byte, length)
// 	_, err = io.ReadFull(conn, data)
// 	if err != nil {
// 		return err
// 	}
// 	err = proto.Unmarshal(data, message)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
