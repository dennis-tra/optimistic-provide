import { Host, Convert as ConvertHost } from "../models/Host";

export class OptProvClient {
  protected baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  public async hostsCreate(): Promise<Host> {
    const resp = await fetch(`${this.baseUrl}/hosts/`, {
      method: "POST",
    });
    const json: Response<Host> = await resp.json();
    if (json.error) {
      throw json.error;
    }
    return ConvertHost.toHost(JSON.stringify(json.data!));
  }

  public async hostGet(hostId: string): Promise<Host> {
    const resp = await fetch(`${this.baseUrl}/hosts/${hostId}`, {
      method: "GET",
    });
    const json: Response<Host> = await resp.json();
    if (json.error) {
      throw json.error;
    }
    return json.data!;
  }

  public async hostsList(): Promise<Host[]> {
    const resp = await fetch(`${this.baseUrl}/hosts/`, {
      method: "GET",
    });
    const json: Response<Host[]> = await resp.json();
    if (json.error) {
      throw json.error;
    }
    return json.data!.map((host) => ConvertHost.toHost(JSON.stringify(host)));
  }

  public async hostBootstrap(hostId: string): Promise<Host> {
    const resp = await fetch(`${this.baseUrl}/hosts/${hostId}/bootstrap`, {
      method: "POST",
    });
    const json: Response<Host> = await resp.json();
    if (json.error) {
      throw json.error;
    }
    return json.data!;
  }

  public async refreshRoutingTable(hostId: string): Promise<void> {
    await fetch(`${this.baseUrl}/hosts/${hostId}/routing-tables/refresh`, {
      method: "POST",
    });
  }

  public async deleteHost(hostId: string): Promise<void> {
    await fetch(`${this.baseUrl}/hosts/${hostId}`, {
      method: "DELETE",
    });
  }
}

export interface Response<T> {
  error: any | null;
  data: T | null;
}
