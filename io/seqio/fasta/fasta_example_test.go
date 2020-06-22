// Copyright ©2020 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fasta_test

import (
	"log"
	"fmt"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
)

func ExampleReader() {
	const multiFasta = `
	>SequenceA
	GATC
	>SequenceB
	CTAG
	>SequenceC
	CTaaaG`

	data := strings.NewReader(multiFasta)
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	r := fasta.NewReader(data, template)
	scanner := seqio.NewScanner(r)

	// Iterate through each sequence in a multifasta
	for scanner.Next() {
		sequence := scanner.Seq()
		asString := sequence.(*linear.Seq).Seq.String()
		asBytes := alphabet.LettersToBytes(sequence.(*linear.Seq).Seq)
		fmt.Printf("%q as %T: %s and %T: %v\n", sequence.Name(), asString, asString, asBytes, asBytes)
	}
	if err := scanner.Error(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// "SequenceA" as string: GATC and []uint8: [71 65 84 67]
	// "SequenceB" as string: CTAG and []uint8: [67 84 65 71]
	// "SequenceC" as string: CTaaaG and []uint8: [67 84 97 97 97 71]
}
