package assetfs

import (
	"regexp"

	"github.com/trustwallet/go-primitives/coin"
)

var (
	regexAssetInfoFile = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/assets/(\w+[\-]\w+|\w+)/info.json$`)
	regexAssetLogoFile = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/assets/(\w+[\-]\w+|\w+)/logo.png$`)

	regexChainInfoFile = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/info/info.json$`)
	regexChainLogoFile = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/info/logo.png$`)

	regexValidatorsAssetLogo = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/validators/assets/(\w+[\-]\w+|\w+)/logo.png$`)
	regexValidatorsList      = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/validators/list.json$`)

	regexDaapLogo = regexp.MustCompile(`./dapps/(\w+[\-]\w+|\w+)/logo.png$`)

	regexTokenListFile = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/tokenlist.json$`)
)

var (
	regexAssetFolder  = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/assets/(\w+[\-]\w+|\w+)$`)
	regexAssetsFolder = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/assets$`)

	regexValidatorsFolder       = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/validators$`)
	regexValidatorsAssetFolder  = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/validators/assets/(\w+[\-]\w+|\w+)$`)
	regexValidatorsAssetsFolder = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/validators/assets$`)

	regexChainFolder     = regexp.MustCompile(`./blockchains/(\w+[^/])$`)
	regexChainInfoFolder = regexp.MustCompile(`./blockchains/(\w+[\-]\w+|\w+)/info$`)
	regexChainsFolder    = regexp.MustCompile(`./blockchains$`)

	regexDappsFolder = regexp.MustCompile(`./dapps$`)
	regexRoot        = regexp.MustCompile(`./$`)
)

var regexes = map[string]*regexp.Regexp{
	TypeAssetInfoFile: regexAssetInfoFile,
	TypeAssetLogoFile: regexAssetLogoFile,

	TypeChainInfoFile: regexChainInfoFile,
	TypeChainLogoFile: regexChainLogoFile,

	TypeValidatorsListFile: regexValidatorsList,
	TypeValidatorsLogoFile: regexValidatorsAssetLogo,

	TypeDappsLogoFile: regexDaapLogo,

	TypeTokenListFile: regexTokenListFile,

	TypeAssetFolder:  regexAssetFolder,
	TypeAssetsFolder: regexAssetsFolder,

	TypeChainFolder:     regexChainFolder,
	TypeChainsFolder:    regexChainsFolder,
	TypeChainInfoFolder: regexChainInfoFolder,

	TypeDaapsFolder: regexDappsFolder,
	TypeRootFolder:  regexRoot,

	TypeValidatorsFolder:       regexValidatorsFolder,
	TypeValidatorsAssetsFolder: regexValidatorsAssetsFolder,
	TypeValidatorsAssetFolder:  regexValidatorsAssetFolder,
}

type Path struct {
	path  string
	chain coin.Coin
	asset string
	type_ string
}

func NewPath(path string) *Path {
	p := Path{
		path: path,
	}

	type_, reg := defineFileType(path)
	if reg == nil {
		p.type_ = TypeUnknown

		return &p
	}

	match := reg.FindStringSubmatch(path)
	if type_ != TypeUnknown {
		p.type_ = type_
	}

	if len(match) >= 2 {
		chain, err := coin.GetCoinForId(match[1])
		if err != nil {
			p.chain = coin.Coin{Handle: match[1]}
		} else {
			p.chain = chain
		}
	}

	if len(match) == 3 {
		p.asset = match[2]
	}

	return &p
}

func (p Path) Type() string {
	return p.type_
}

func (p Path) String() string {
	return p.path
}

func (p Path) Chain() coin.Coin {
	return p.chain
}

func (p Path) Asset() string {
	return p.asset
}

func defineFileType(p string) (string, *regexp.Regexp) {
	for t, r := range regexes {
		if r.MatchString(p) {
			return t, r
		}
	}

	return TypeUnknown, nil
}
