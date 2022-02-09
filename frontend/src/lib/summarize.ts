type SummarizeOptions = {
  mode: 'abs' | 'ext'
  type: 'txt' | 'url' | 'pdf'
  input: string
  maxChunk?: number
  maxCharPerSlide?: number
}

type SummarizeResponse = {
  fileName: string
  wordsAfter: number
  wordsBefore: number
  reducedByPercentage: number
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
    return (await (
      await fetch(`${this.ServerURL}/summarize`, {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(opts),
      })
    ).json()) as SummarizeResponse
  }

  public summarize = async (opts: SummarizeOptions): Promise<string> => {
    switch (opts.type) {
      case 'url':
        const res = await this.summarizeURL(opts)
        return await this.getSlideMDText(res.fileName)
      case 'txt':
        throw new Error('not implemented')
      case 'pdf':
        throw new Error('not implemented')
    }
  }
}
