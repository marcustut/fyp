import { Instance } from '@/generated/graphql'

// ComputeOptions is the options that can be
// provided in a request to the Compute API.
type ComputeOptions = {
  slide_id: string
}

// config to setup ComputeClient
type ComputeClientConfig = {
  ServerURL: string
}

// type of request that can be made to the
// Compute API.
type RequestType = 'run' | 'terminate'

// dyanmic return type based on RequestType
type SendRequestReturnType<T> = T extends 'run'
  ? Omit<Instance, '__typename'> | string
  : T extends 'terminate'
  ? Omit<Instance, '__typename'> | string
  : never

// ComputeClient is an API client to interact with the
// Compute API.
export class ComputeClient {
  ServerURL: string

  constructor(config: ComputeClientConfig) {
    this.ServerURL = config.ServerURL
  }

  /**
   * sendRequest is a private method that is used internally to
   * handle sending http requests to the Compute API.
   * @param {RequestType} type - type of request made to the API.
   * @param {ComputeOptions} opts - request options for the API call.
   * @returns
   */
  private async sendRequest<T extends RequestType>(
    type: T,
    opts: ComputeOptions,
    method: 'POST' | 'GET'
  ): Promise<SendRequestReturnType<T>> {
    return await (
      await fetch(`${this.ServerURL}/${type}`, {
        method,
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(opts),
      })
    ).json()
  }

  /**
   * runInstance runs a single EC2 instance.
   * @param {ComputeOptions} opts - request options for the API call.
   * @returns
   */
  public async runInstance(opts: ComputeOptions) {
    return await this.sendRequest('run', opts, 'POST')
  }

  /**
   * terminateInstances terminates all EC2 instance for a given slide_id.
   * @param {ComputeOptions} opts - request options for the API call.
   * @returns
   */
  public async terminateInstances(opts: ComputeOptions) {
    return await this.sendRequest('terminate', opts, 'POST')
  }
}
