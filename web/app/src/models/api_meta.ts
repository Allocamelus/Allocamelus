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

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_MetaCaptchaSiteKeys(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
    this.siteKeys = source["site-keys"]
    this.difficulties = source["difficulties"]
  }

  siteKey(key: "easy" | "moderate" | "hard" | "all" | string) {
    var siteKey: string
    try {
      siteKey = this.siteKeys[key]
    } catch {
      siteKey = this.siteKeys.all
    }
    return siteKey
  }
}
