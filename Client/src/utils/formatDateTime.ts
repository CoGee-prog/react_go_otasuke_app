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

const weekDays = ['日', '月', '火', '水', '木', '金', '土'];

export const formatTimeRange = (startTime: string, endTime: string) => {
  const start = new Date(startTime);
  const end = new Date(endTime);

  const formatDate = (date: Date) => {
    const dayOfWeek = weekDays[date.getDay()];
    let color = '';
    if (dayOfWeek === '土') {
      color = 'green';
    } else if (dayOfWeek === '日') {
      color = 'red';
    }
    return {
      dateText: `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`,
      dayOfWeekText: `(${dayOfWeek})`,
      dayOfWeekColor: color,
    };
  };

  const formatTime = (date: Date) => {
    return `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`;
  };

  const startDateFormatted = formatDate(start);
  const endDateFormatted = formatDate(end);

  if (
    start.getFullYear() === end.getFullYear() &&
    start.getMonth() === end.getMonth() &&
    start.getDate() === end.getDate()
  ) {
    // 同じ日の場合、開始の日付と開始・終了の時間を表示
    return {
      text: `${startDateFormatted.dateText} ${startDateFormatted.dayOfWeekText} ${formatTime(start)}~${formatTime(end)}`,
      dayOfWeekColor: startDateFormatted.dayOfWeekColor,
    };
  } else {
    // 異なる日の場合、開始と終了の日付と時間を表示
    return {
      text: `${startDateFormatted.dateText} ${startDateFormatted.dayOfWeekText} ${formatTime(start)} ~ ${endDateFormatted.dateText} ${endDateFormatted.dayOfWeekText} ${formatTime(end)}`,
      startDayOfWeekColor: startDateFormatted.dayOfWeekColor,
      endDayOfWeekColor: endDateFormatted.dayOfWeekColor,
    };
  }
};


export default formatDateTime;