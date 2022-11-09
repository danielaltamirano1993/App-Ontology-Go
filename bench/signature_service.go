package bench

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/signature"
)

func (this *TestTransfer) SignatureService() {
	now := time.Now().UnixNano()
	s := strconv.Itoa(int(now))
	buf := []byte(s)
	temp := sha256.Sum256(buf)
	hash := common.Uint256(sha256.Sum256(temp[:]))
	log.Infof("hash:%x", hash)

	schemes := []string{
		"SHA224withECDSA",
		"SHA256withECDSA",
		"SHA384withECDSA",
		"SHA512withECDSA",
		"SHA3-224withECDSA",
		"SHA3-256withECDSA",
		"SHA3-384withECDSA",
		"SHA3-512withECDSA",
		"RIPEMD160withECDSA",
		"SM3withSM2",
		"SHA512withEdDSA",
	}

	var signers []*account.Account
	for _, scheme := range schemes {
		signer := account.NewAccount(scheme)
		signers = append(signers, signer)
	}
	for _, signer := range signers {
		sig, err := signature.Sign(signer, hash[:])
		if err != nil {
			log.Errorf("%d sign data err:%s", signer.Scheme(), err)
			continue
		}
		err = signature.Verify(signer.PublicKey, hash[:], sig)
		if err != nil {
			log.Errorf("verify error:%s, scheme:%d", err, signer.Scheme())
			continue
		}
		log.Infof("check sign :%d pass", signer.Scheme())
	}

	rand.Seed(time.Now().UnixNano())
	M := rand.Intn(len(signers))
	dst := make([]*account.Account, len(signers))
	perm := rand.Perm(len(signers))
	for i, v := range perm {
		dst[v] = signers[i]
	}
	log.Infof("M is :%d", M)

	var sigs [][]byte
	var pks []keypair.PublicKey
	for i := 0; i < (M + 1); i++ {
		signer := dst[i]
		pks = append(pks, signer.PublicKey)
		log.Infof("signer:%x %d", keypair.SerializePublicKey(signer.PublicKey), signer.Scheme())
		sig, err := signature.Sign(signer, hash[:])
		if err != nil {
			log.Errorf("multi sig failed:%s", err)
			return
		}
		sigs = append(sigs, sig)
	}
	err := signature.VerifyMultiSignature(hash[:], pks, M, sigs)
	if err != nil {
		log.Errorf(" multi sig verify failed :%s", err)
	}
	log.Infof("multi sig verify pass")

}
