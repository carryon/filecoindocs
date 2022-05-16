package types

import (
	"fmt"
	"io"
	"math"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *BlockHeader) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{140}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Ticket (types.Ticket) (struct)
	if err := t.Ticket.MarshalCBOR(w); err != nil {
		return err
	}

	// t.EPostProof (types.EPostProof) (struct)
	if err := t.EPostProof.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Parents ([]cid.Cid) (slice)
	if len(t.Parents) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Parents was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Parents)))); err != nil {
		return err
	}
	for _, v := range t.Parents {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Parents: %w", err)
		}
	}

	// t.ParentWeight (types.BigInt) (struct)
	if err := t.ParentWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Height (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Height))); err != nil {
		return err
	}

	// t.ParentStateRoot (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.ParentStateRoot); err != nil {
		return xerrors.Errorf("failed to write cid field t.ParentStateRoot: %w", err)
	}

	// t.ParentMessageReceipts (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.ParentMessageReceipts); err != nil {
		return xerrors.Errorf("failed to write cid field t.ParentMessageReceipts: %w", err)
	}

	// t.Messages (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Messages); err != nil {
		return xerrors.Errorf("failed to write cid field t.Messages: %w", err)
	}

	// t.BLSAggregate (types.Signature) (struct)
	if err := t.BLSAggregate.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Timestamp (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Timestamp))); err != nil {
		return err
	}

	// t.BlockSig (types.Signature) (struct)
	if err := t.BlockSig.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *BlockHeader) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 12 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Ticket (types.Ticket) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Ticket = new(Ticket)
			if err := t.Ticket.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.EPostProof (types.EPostProof) (struct)

	{

		if err := t.EPostProof.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Parents ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Parents: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Parents = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Parents failed: %w", err)
		}
		t.Parents[i] = c
	}

	// t.ParentWeight (types.BigInt) (struct)

	{

		if err := t.ParentWeight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Height (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Height = uint64(extra)
	// t.ParentStateRoot (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ParentStateRoot: %w", err)
		}

		t.ParentStateRoot = c

	}
	// t.ParentMessageReceipts (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ParentMessageReceipts: %w", err)
		}

		t.ParentMessageReceipts = c

	}
	// t.Messages (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Messages: %w", err)
		}

		t.Messages = c

	}
	// t.BLSAggregate (types.Signature) (struct)

	{

		if err := t.BLSAggregate.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Timestamp (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Timestamp = uint64(extra)
	// t.BlockSig (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.BlockSig = new(Signature)
			if err := t.BlockSig.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *Ticket) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.VRFProof ([]uint8) (slice)
	if len(t.VRFProof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.VRFProof was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.VRFProof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.VRFProof); err != nil {
		return err
	}
	return nil
}

func (t *Ticket) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.VRFProof ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.VRFProof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.VRFProof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.VRFProof); err != nil {
		return err
	}
	return nil
}

func (t *EPostProof) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Proof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Proof); err != nil {
		return err
	}

	// t.PostRand ([]uint8) (slice)
	if len(t.PostRand) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.PostRand was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.PostRand)))); err != nil {
		return err
	}
	if _, err := w.Write(t.PostRand); err != nil {
		return err
	}

	// t.Candidates ([]types.EPostTicket) (slice)
	if len(t.Candidates) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Candidates was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Candidates)))); err != nil {
		return err
	}
	for _, v := range t.Candidates {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *EPostProof) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proof ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Proof = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Proof); err != nil {
		return err
	}
	// t.PostRand ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.PostRand: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.PostRand = make([]byte, extra)
	if _, err := io.ReadFull(br, t.PostRand); err != nil {
		return err
	}
	// t.Candidates ([]types.EPostTicket) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Candidates: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Candidates = make([]EPostTicket, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v EPostTicket
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Candidates[i] = v
	}

	return nil
}

func (t *EPostTicket) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Partial ([]uint8) (slice)
	if len(t.Partial) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Partial was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Partial)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Partial); err != nil {
		return err
	}

	// t.SectorID (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorID))); err != nil {
		return err
	}

	// t.ChallengeIndex (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ChallengeIndex))); err != nil {
		return err
	}
	return nil
}

func (t *EPostTicket) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Partial ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Partial: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Partial = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Partial); err != nil {
		return err
	}
	// t.SectorID (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SectorID = uint64(extra)
	// t.ChallengeIndex (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.ChallengeIndex = uint64(extra)
	return nil
}

func (t *Message) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{136}); err != nil {
		return err
	}

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.From (address.Address) (struct)
	if err := t.From.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}

	// t.Value (types.BigInt) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}

	// t.GasPrice (types.BigInt) (struct)
	if err := t.GasPrice.MarshalCBOR(w); err != nil {
		return err
	}

	// t.GasLimit (types.BigInt) (struct)
	if err := t.GasLimit.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Method (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Method))); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Params)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Params); err != nil {
		return err
	}
	return nil
}

func (t *Message) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.From (address.Address) (struct)

	{

		if err := t.From.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	// t.Value (types.BigInt) (struct)

	{

		if err := t.Value.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.GasPrice (types.BigInt) (struct)

	{

		if err := t.GasPrice.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.GasLimit (types.BigInt) (struct)

	{

		if err := t.GasLimit.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Method (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Method = uint64(extra)
	// t.Params ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Params = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Params); err != nil {
		return err
	}
	return nil
}

func (t *SignedMessage) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Message (types.Message) (struct)
	if err := t.Message.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedMessage) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Message (types.Message) (struct)

	{

		if err := t.Message.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Signature (types.Signature) (struct)

	{

		if err := t.Signature.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *MsgMeta) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.BlsMessages (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.BlsMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.BlsMessages: %w", err)
	}

	// t.SecpkMessages (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.SecpkMessages); err != nil {
		return xerrors.Errorf("failed to write cid field t.SecpkMessages: %w", err)
	}

	return nil
}

func (t *MsgMeta) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.BlsMessages (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.BlsMessages: %w", err)
		}

		t.BlsMessages = c

	}
	// t.SecpkMessages (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SecpkMessages: %w", err)
		}

		t.SecpkMessages = c

	}
	return nil
}

func (t *SignedVoucher) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{137}); err != nil {
		return err
	}

	// t.TimeLock (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.TimeLock))); err != nil {
		return err
	}

	// t.SecretPreimage ([]uint8) (slice)
	if len(t.SecretPreimage) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.SecretPreimage was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.SecretPreimage)))); err != nil {
		return err
	}
	if _, err := w.Write(t.SecretPreimage); err != nil {
		return err
	}

	// t.Extra (types.ModVerifyParams) (struct)
	if err := t.Extra.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Lane))); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}

	// t.Amount (types.BigInt) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinCloseHeight (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MinCloseHeight))); err != nil {
		return err
	}

	// t.Merges ([]types.Merge) (slice)
	if len(t.Merges) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Merges was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Merges)))); err != nil {
		return err
	}
	for _, v := range t.Merges {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedVoucher) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.TimeLock (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.TimeLock = uint64(extra)
	// t.SecretPreimage ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.SecretPreimage: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.SecretPreimage = make([]byte, extra)
	if _, err := io.ReadFull(br, t.SecretPreimage); err != nil {
		return err
	}
	// t.Extra (types.ModVerifyParams) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Extra = new(ModVerifyParams)
			if err := t.Extra.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Lane (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Lane = uint64(extra)
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	// t.Amount (types.BigInt) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinCloseHeight (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.MinCloseHeight = uint64(extra)
	// t.Merges ([]types.Merge) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Merges: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Merges = make([]Merge, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v Merge
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Merges[i] = v
	}

	// t.Signature (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *ModVerifyParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Actor (address.Address) (struct)
	if err := t.Actor.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Method (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Method))); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Data)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Data); err != nil {
		return err
	}
	return nil
}

func (t *ModVerifyParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Actor (address.Address) (struct)

	{

		if err := t.Actor.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Method (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Method = uint64(extra)
	// t.Data ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Data = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Data); err != nil {
		return err
	}
	return nil
}

func (t *Merge) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Lane))); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}
	return nil
}

func (t *Merge) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Lane (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Lane = uint64(extra)
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	return nil
}

func (t *Actor) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.Code (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Code); err != nil {
		return xerrors.Errorf("failed to write cid field t.Code: %w", err)
	}

	// t.Head (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Head); err != nil {
		return xerrors.Errorf("failed to write cid field t.Head: %w", err)
	}

	// t.Nonce (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}

	// t.Balance (types.BigInt) (struct)
	if err := t.Balance.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Actor) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Code (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Code: %w", err)
		}

		t.Code = c

	}
	// t.Head (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Head: %w", err)
		}

		t.Head = c

	}
	// t.Nonce (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Nonce = uint64(extra)
	// t.Balance (types.BigInt) (struct)

	{

		if err := t.Balance.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *MessageReceipt) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.ExitCode (uint8) (uint8)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ExitCode))); err != nil {
		return err
	}

	// t.Return ([]uint8) (slice)
	if len(t.Return) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Return was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Return)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Return); err != nil {
		return err
	}

	// t.GasUsed (types.BigInt) (struct)
	if err := t.GasUsed.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *MessageReceipt) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ExitCode (uint8) (uint8)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.ExitCode = uint8(extra)
	// t.Return ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Return: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Return = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Return); err != nil {
		return err
	}
	// t.GasUsed (types.BigInt) (struct)

	{

		if err := t.GasUsed.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *BlockMsg) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Header (types.BlockHeader) (struct)
	if err := t.Header.MarshalCBOR(w); err != nil {
		return err
	}

	// t.BlsMessages ([]cid.Cid) (slice)
	if len(t.BlsMessages) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.BlsMessages was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.BlsMessages)))); err != nil {
		return err
	}
	for _, v := range t.BlsMessages {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.BlsMessages: %w", err)
		}
	}

	// t.SecpkMessages ([]cid.Cid) (slice)
	if len(t.SecpkMessages) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.SecpkMessages was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.SecpkMessages)))); err != nil {
		return err
	}
	for _, v := range t.SecpkMessages {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.SecpkMessages: %w", err)
		}
	}
	return nil
}

func (t *BlockMsg) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Header (types.BlockHeader) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Header = new(BlockHeader)
			if err := t.Header.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.BlsMessages ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BlsMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.BlsMessages = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.BlsMessages failed: %w", err)
		}
		t.BlsMessages[i] = c
	}

	// t.SecpkMessages ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.SecpkMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.SecpkMessages = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.SecpkMessages failed: %w", err)
		}
		t.SecpkMessages[i] = c
	}

	return nil
}

func (t *SignedStorageAsk) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Ask (types.StorageAsk) (struct)
	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedStorageAsk) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Ask (types.StorageAsk) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Ask = new(StorageAsk)
			if err := t.Ask.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Signature (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *StorageAsk) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{134}); err != nil {
		return err
	}

	// t.Price (types.BigInt) (struct)
	if err := t.Price.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinPieceSize (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MinPieceSize))); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Timestamp (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Timestamp))); err != nil {
		return err
	}

	// t.Expiry (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Expiry))); err != nil {
		return err
	}

	// t.SeqNo (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SeqNo))); err != nil {
		return err
	}
	return nil
}

func (t *StorageAsk) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Price (types.BigInt) (struct)

	{

		if err := t.Price.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinPieceSize (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.MinPieceSize = uint64(extra)
	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Timestamp (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Timestamp = uint64(extra)
	// t.Expiry (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Expiry = uint64(extra)
	// t.SeqNo (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SeqNo = uint64(extra)
	return nil
}

func (t *ExpTipSet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Cids ([]cid.Cid) (slice)
	if len(t.Cids) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Cids was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Cids)))); err != nil {
		return err
	}
	for _, v := range t.Cids {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Cids: %w", err)
		}
	}

	// t.Blocks ([]*types.BlockHeader) (slice)
	if len(t.Blocks) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Blocks was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Blocks)))); err != nil {
		return err
	}
	for _, v := range t.Blocks {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Height (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Height))); err != nil {
		return err
	}
	return nil
}

func (t *ExpTipSet) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Cids ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Cids: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Cids = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Cids failed: %w", err)
		}
		t.Cids[i] = c
	}

	// t.Blocks ([]*types.BlockHeader) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Blocks: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Blocks = make([]*BlockHeader, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v BlockHeader
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Blocks[i] = &v
	}

	// t.Height (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Height = uint64(extra)
	return nil
}