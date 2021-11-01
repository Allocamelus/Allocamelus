export declare class Go {
  constructor();
  importObject: any;
  run(instance: WebAssembly.Instance);
}
export interface Go {
  constructor();
  importObject: any;
  run(instance: WebAssembly.Instance);
}
declare global {
  interface Window {
    Go?: Go & typeof Go;
  }
}
