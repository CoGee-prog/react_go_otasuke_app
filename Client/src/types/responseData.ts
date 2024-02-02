export interface ResponseData<T> {
  status: number;
	message: string;
	result: T;
}