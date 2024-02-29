// APIからのレスポンスの日時をフォーマットする
function formatDateTime(isoString: string): string {
  const date = new Date(isoString);
  return new Intl.DateTimeFormat('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
		weekday: 'short',
    hour: '2-digit',
    minute: '2-digit',
  }).format(date);
}

export default formatDateTime;