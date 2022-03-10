export enum SummarizeMode {
  Abstractive = 'abs',
  Extractive = 'ext'
}

export enum SummarizeType {
  Text = 'txt',
  URL = 'url',
  PDF = 'pdf'
}

export enum SummarizeMaxChunk {
  Long = 150,
  Intermediate = 250,
  Short = 500
}

export enum SummarizeMaxCharPerSlide {
  Compact = 500, 
  Comfortable = 250
}

export enum SummarizeTheme {
  Default = 'default',
  AppleBasic = 'apple-basic',
  Seriph = 'seriph',
  Bricks = 'bricks',
  ShibaInu = 'shibainu'
}

export type SummarizeOptions = {
  mode: SummarizeMode
  type: SummarizeType
  input: string
  outputName: string
  maxChunk?: SummarizeMaxChunk
  maxCharPerSlide?: SummarizeMaxCharPerSlide
  theme?: SummarizeTheme
}

type SummarizeResponse = {
  data?: {
    fileName: string
    wordsAfter: number
    wordsBefore: number
    reducedByPercentage: number
  }
  error?: Error
}

export type SummarizeEstimateResponse = {
  articleLength: number; 
  minutes: number;
  seconds: number;
}

// config to setup SummarizeClient
type SummarizeClientConfig = {
  ServerURL: string
}

// SummarizeClient is an API client to interact with the Summarize API.
export class SummarizeClient {
  ServerURL: string

  constructor(config: SummarizeClientConfig) {
    this.ServerURL = config.ServerURL
  }

  /**
   * getSlideMDText fetches the text content of summarized markdown file.
   * @param {string} filename - filename of the slide generated on server-side, eg. slides.txt.
   * @returns {string} text string of the summarized markdown file.
   */
  private getSlideMDText = async (filename: string): Promise<string> =>
    await (await fetch(`${this.ServerURL}/output/${filename}`)).text()

  /**
   * summarizeURL takes a URL and send to the Summarize API for
   * scraping the page content and summarize it.
   * @param {SummarizeOptions} opts - options for summarizing the text.
   * @returns {SummarizeResponse} response returned by the API.
   */
  private summarizeURL = async (
    opts: SummarizeOptions
  ): Promise<SummarizeResponse> => {
    const res = await fetch(`${this.ServerURL}/summarize`, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(opts),
    })
    if (!res.ok) return { error: new Error('failed to summarize') }
    return { data: await res.json() }
  }

  public summarize = async (opts: SummarizeOptions): Promise<string> => {
    const { data, error } = await this.summarizeURL(opts)
    if (error || !data) {
      throw error
    }
    return await this.getSlideMDText(data.fileName)
  }

  public uploadFile = async (
    blob: Blob,
    fileName: string
  ): Promise<{ message: string }> => {
    const formData = new FormData()

    formData.append('file', blob, fileName)

    const response = await fetch(`${this.ServerURL}/uploads`, {
      method: 'POST',
      body: formData,
    })
    const resp = await response.json()

    return resp
  }

  public estimate = async (
    type: SummarizeType,
    input: string
  ): Promise<SummarizeEstimateResponse> => {
    const response = await fetch(`${this.ServerURL}/estimate?type=${type}&input=${input}`, {
      method: 'POST',
    })
    return await response.json()
  }
}
