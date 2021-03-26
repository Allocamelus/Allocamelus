export class API_MetaCaptchaSiteKeys {
  siteKeys: {
    easy: string,
    moderate: string,
    hard: string,
    all: string,
  };
  difficulties: {
    user: {
      create: string
    }
  }

  static createFrom(source: any = {}) {
    return new API_MetaCaptchaSiteKeys(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.siteKeys = source["site-keys"]
    this.difficulties = source["difficulties"]
  }

  siteKey(key: "easy" | "moderate" | "hard" | "all" | string) {
    var siteKey
    try {
      siteKey = this.siteKeys[key]
    } catch {
      siteKey = this.siteKeys.all
    }
    return siteKey
  }
}
