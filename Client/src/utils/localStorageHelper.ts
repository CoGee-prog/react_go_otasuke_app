export const saveDataWithExpiry = <T>(key: string, data: T, expiryInMinutes: number): void => {
  const now = new Date();
	 // 分単位で有効期限を設定
  const expiryTime = new Date(now.getTime() + expiryInMinutes * 60000);

  const dataWithExpiry = {
    value: data,
    expiry: expiryTime.getTime(),
  };

  localStorage.setItem(key, JSON.stringify(dataWithExpiry));
};

export const loadDataWithExpiry = <T>(key: string): T | null => {
  if (typeof localStorage === 'undefined') {
    return null;
  }
	
  const dataWithExpiryString = localStorage.getItem(key);

  if (!dataWithExpiryString) {
    return null;
  }

  const dataWithExpiry = JSON.parse(dataWithExpiryString) as { value: T; expiry: number };

  const now = new Date();
  if (now.getTime() > dataWithExpiry.expiry) {
    localStorage.removeItem(key);
    return null;
  }

  return dataWithExpiry.value;
};
