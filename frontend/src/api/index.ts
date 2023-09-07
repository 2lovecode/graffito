async function post<T, A>(
    url: string,
    data: A
): Promise<T> {
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });
  
      if (!response.ok) {
        throw new Error(`Request failed with status ${response.status}`);
      }
  
      const responseData: T = await response.json();
      return responseData;
    } catch (error) {
      throw new Error(`Request failed: ${error}`);
    }
}

interface SandboxRunJson {
    sourceCode: string
}
interface SandboxRunData {
    content: string
}
  // 示例数据接口
interface SandboxRunResp {
    status: string
    errorCode: string
    errorMessage: string
    data: SandboxRunData
}

  // 示例用法：
  export async function sandboxRun(data: SandboxRunJson):Promise<SandboxRunData> {
    try {
      const url = 'http://localhost:8081/apps/sandbox/run';
      const postData = await post<SandboxRunResp, SandboxRunJson>(url, data);
  
      return postData.data
    } catch (error) {
      // 在发生错误时应该进行适当的错误处理
        throw new Error(`Sandbox Run failed: ${error}`);
    }
  }
  