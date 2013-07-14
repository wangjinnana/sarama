package kafka

type encoder interface {
	encode(pe packetEncoder)
}

func encode(in encoder) ([]byte, error) {
	if in == nil {
		return nil, nil
	}

	var prepEnc prepEncoder
	var realEnc realEncoder

	in.encode(&prepEnc)
	if prepEnc.err != nil {
		return nil, prepEnc.err
	}

	realEnc.raw = make([]byte, prepEnc.length)
	in.encode(&realEnc)

	return realEnc.raw, nil
}

type decoder interface {
	decode(pd packetDecoder) error
}

func decode(buf []byte, in decoder) error {
	if buf == nil {
		return nil
	}

	helper := realDecoder{raw: buf}
	return in.decode(&helper)
}

type encoderDecoder interface {
	encoder
	decoder
}
