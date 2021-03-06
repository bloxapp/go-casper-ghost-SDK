// Code generated by fastssz. DO NOT EDIT.
package core

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the PendingAttestation object
func (p *PendingAttestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PendingAttestation object to a target array
func (p *PendingAttestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(148)

	// Offset (0) 'AggregationBits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(p.AggregationBits)

	// Field (1) 'Data'
	if p.Data == nil {
		p.Data = new(AttestationData)
	}
	if dst, err = p.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'InclusionDelay'
	dst = ssz.MarshalUint64(dst, p.InclusionDelay)

	// Field (3) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, p.ProposerIndex)

	// Field (0) 'AggregationBits'
	if len(p.AggregationBits) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, p.AggregationBits...)

	return
}

// UnmarshalSSZ ssz unmarshals the PendingAttestation object
func (p *PendingAttestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 148 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AggregationBits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if p.Data == nil {
		p.Data = new(AttestationData)
	}
	if err = p.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'InclusionDelay'
	p.InclusionDelay = ssz.UnmarshallUint64(buf[132:140])

	// Field (3) 'ProposerIndex'
	p.ProposerIndex = ssz.UnmarshallUint64(buf[140:148])

	// Field (0) 'AggregationBits'
	{
		buf = tail[o0:]
		if err = ssz.ValidateBitlist(buf, 2048); err != nil {
			return err
		}
		if cap(p.AggregationBits) == 0 {
			p.AggregationBits = make([]byte, 0, len(buf))
		}
		p.AggregationBits = append(p.AggregationBits, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PendingAttestation object
func (p *PendingAttestation) SizeSSZ() (size int) {
	size = 148

	// Field (0) 'AggregationBits'
	size += len(p.AggregationBits)

	return
}

// HashTreeRoot ssz hashes the PendingAttestation object
func (p *PendingAttestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PendingAttestation object with a hasher
func (p *PendingAttestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AggregationBits'
	if len(p.AggregationBits) == 0 {
		err = ssz.ErrEmptyBitlist
		return
	}
	hh.PutBitlist(p.AggregationBits, 2048)

	// Field (1) 'Data'
	if err = p.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'InclusionDelay'
	hh.PutUint64(p.InclusionDelay)

	// Field (3) 'ProposerIndex'
	hh.PutUint64(p.ProposerIndex)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the Attestation object
func (a *Attestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the Attestation object to a target array
func (a *Attestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(228)

	// Offset (0) 'AggregationBits'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(a.AggregationBits)

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if dst, err = a.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(a.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, a.Signature...)

	// Field (0) 'AggregationBits'
	if len(a.AggregationBits) > 2048 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, a.AggregationBits...)

	return
}

// UnmarshalSSZ ssz unmarshals the Attestation object
func (a *Attestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AggregationBits'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if a.Data == nil {
		a.Data = new(AttestationData)
	}
	if err = a.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	if cap(a.Signature) == 0 {
		a.Signature = make([]byte, 0, len(buf[132:228]))
	}
	a.Signature = append(a.Signature, buf[132:228]...)

	// Field (0) 'AggregationBits'
	{
		buf = tail[o0:]
		if err = ssz.ValidateBitlist(buf, 2048); err != nil {
			return err
		}
		if cap(a.AggregationBits) == 0 {
			a.AggregationBits = make([]byte, 0, len(buf))
		}
		a.AggregationBits = append(a.AggregationBits, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Attestation object
func (a *Attestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AggregationBits'
	size += len(a.AggregationBits)

	return
}

// HashTreeRoot ssz hashes the Attestation object
func (a *Attestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the Attestation object with a hasher
func (a *Attestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AggregationBits'
	if len(a.AggregationBits) == 0 {
		err = ssz.ErrEmptyBitlist
		return
	}
	hh.PutBitlist(a.AggregationBits, 2048)

	// Field (1) 'Data'
	if err = a.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(a.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(a.Signature)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the AttestationData object
func (a *AttestationData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AttestationData object to a target array
func (a *AttestationData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, a.Slot)

	// Field (1) 'CommitteeIndex'
	dst = ssz.MarshalUint64(dst, a.CommitteeIndex)

	// Field (2) 'BeaconBlockRoot'
	if len(a.BeaconBlockRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, a.BeaconBlockRoot...)

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if dst, err = a.Source.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if dst, err = a.Target.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AttestationData object
func (a *AttestationData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 128 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	a.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'CommitteeIndex'
	a.CommitteeIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'BeaconBlockRoot'
	if cap(a.BeaconBlockRoot) == 0 {
		a.BeaconBlockRoot = make([]byte, 0, len(buf[16:48]))
	}
	a.BeaconBlockRoot = append(a.BeaconBlockRoot, buf[16:48]...)

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if err = a.Source.UnmarshalSSZ(buf[48:88]); err != nil {
		return err
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if err = a.Target.UnmarshalSSZ(buf[88:128]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AttestationData object
func (a *AttestationData) SizeSSZ() (size int) {
	size = 128
	return
}

// HashTreeRoot ssz hashes the AttestationData object
func (a *AttestationData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AttestationData object with a hasher
func (a *AttestationData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(a.Slot)

	// Field (1) 'CommitteeIndex'
	hh.PutUint64(a.CommitteeIndex)

	// Field (2) 'BeaconBlockRoot'
	if len(a.BeaconBlockRoot) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(a.BeaconBlockRoot)

	// Field (3) 'Source'
	if err = a.Source.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'Target'
	if err = a.Target.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the IndexedAttestation object
func (i *IndexedAttestation) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the IndexedAttestation object to a target array
func (i *IndexedAttestation) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(228)

	// Offset (0) 'AttestingIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(i.AttestingIndices) * 8

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if dst, err = i.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(i.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, i.Signature...)

	// Field (0) 'AttestingIndices'
	if len(i.AttestingIndices) > 2048 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(i.AttestingIndices); ii++ {
		dst = ssz.MarshalUint64(dst, i.AttestingIndices[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the IndexedAttestation object
func (i *IndexedAttestation) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 228 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'AttestingIndices'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'Data'
	if i.Data == nil {
		i.Data = new(AttestationData)
	}
	if err = i.Data.UnmarshalSSZ(buf[4:132]); err != nil {
		return err
	}

	// Field (2) 'Signature'
	if cap(i.Signature) == 0 {
		i.Signature = make([]byte, 0, len(buf[132:228]))
	}
	i.Signature = append(i.Signature, buf[132:228]...)

	// Field (0) 'AttestingIndices'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 8, 2048)
		if err != nil {
			return err
		}
		i.AttestingIndices = ssz.ExtendUint64(i.AttestingIndices, num)
		for ii := 0; ii < num; ii++ {
			i.AttestingIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the IndexedAttestation object
func (i *IndexedAttestation) SizeSSZ() (size int) {
	size = 228

	// Field (0) 'AttestingIndices'
	size += len(i.AttestingIndices) * 8

	return
}

// HashTreeRoot ssz hashes the IndexedAttestation object
func (i *IndexedAttestation) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the IndexedAttestation object with a hasher
func (i *IndexedAttestation) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'AttestingIndices'
	{
		if len(i.AttestingIndices) > 2048 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range i.AttestingIndices {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(i.AttestingIndices))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(2048, numItems, 8))
	}

	// Field (1) 'Data'
	if err = i.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	if len(i.Signature) != 96 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(i.Signature)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the Checkpoint object
func (c *Checkpoint) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Checkpoint object to a target array
func (c *Checkpoint) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Epoch'
	dst = ssz.MarshalUint64(dst, c.Epoch)

	// Field (1) 'Root'
	if len(c.Root) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, c.Root...)

	return
}

// UnmarshalSSZ ssz unmarshals the Checkpoint object
func (c *Checkpoint) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 40 {
		return ssz.ErrSize
	}

	// Field (0) 'Epoch'
	c.Epoch = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Root'
	if cap(c.Root) == 0 {
		c.Root = make([]byte, 0, len(buf[8:40]))
	}
	c.Root = append(c.Root, buf[8:40]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Checkpoint object
func (c *Checkpoint) SizeSSZ() (size int) {
	size = 40
	return
}

// HashTreeRoot ssz hashes the Checkpoint object
func (c *Checkpoint) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Checkpoint object with a hasher
func (c *Checkpoint) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Epoch'
	hh.PutUint64(c.Epoch)

	// Field (1) 'Root'
	if len(c.Root) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(c.Root)

	hh.Merkleize(indx)
	return
}
