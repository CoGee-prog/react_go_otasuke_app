import * as React from 'react'
import { Box, TextField } from '@mui/material'
import {
  LocalizationProvider,
  DatePicker,
  PickersDayProps,
  PickersDay,
} from '@mui/x-date-pickers-pro'
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns'
import ja from 'date-fns/locale/ja'

interface CustomDatePickerProps {
  value: Date | null
  onChange: (newValue: Date | null) => void
  error?: boolean
  helperText?: string
}

const CustomDatePicker: React.FC<CustomDatePickerProps> = ({
  value,
  onChange,
  error,
  helperText,
}) => {
  const styles = {
    mobiledialogprops: {
      '.MuiDatePickerToolbar-title': {
        fontSize: '1.5rem',
      },
      '.MuiDayPicker-header span:nth-of-type(1)': {
        color: 'rgba(255, 0, 0, 0.6)',
      },
      '.MuiDayPicker-header span:nth-of-type(7)': {
        color: 'rgba(0, 0, 255, 0.6)',
      },
    },
    paperprops: {
      '.MuiDayPicker-header span:nth-of-type(1)': {
        color: 'rgba(255, 0, 0, 0.6)',
      },
      '.MuiDayPicker-header span:nth-of-type(7)': {
        color: 'rgba(0, 0, 255, 0.6)',
      },
    },
  }
  const renderWeekEndPickerDay = (
    date: Date,
    _selectedDates: Array<Date | null>,
    pickersDayProps: PickersDayProps<Date>,
  ) => {
    const switchDayColor = (getday: number) => {
      switch (getday) {
        case 0:
          return { color: 'red' }
        case 6:
          return { color: 'blue' }
        default:
          return {}
      }
    }
    const newPickersDayProps = {
      ...pickersDayProps,
      sx: switchDayColor(date.getDay()),
    }
    return <PickersDay {...newPickersDayProps} />
  }
  return (
    <LocalizationProvider
      dateAdapter={AdapterDateFns}
      adapterLocale={ja}
      dateFormats={{ monthAndYear: 'yyyy年 MM月', year: 'yyyy年' }}
      localeText={{
        previousMonth: '前月を表示',
        nextMonth: '次月を表示',
        cancelButtonLabel: 'キャンセル',
        okButtonLabel: '選択',
      }}
    >
      <Box
        sx={{
          borderRadius: 2,
          marginLeft: 'auto',
          marginRight: 'auto',
        }}
      >
        <DatePicker
          label='日付'
          minDate={new Date(new Date().getTime() + 9 * 60 * 60 * 1000)}
          value={value}
          onChange={onChange}
          inputFormat='yyyy/MM/dd'
          mask='____年__月__日'
          toolbarFormat='yyyy年MM月dd日'
          renderInput={(params) => (
            <TextField
              {...params}
              fullWidth
              inputProps={{
                ...params.inputProps,
                placeholder: '年/月/日',
              }}
              error={error}
              helperText={helperText}
            />
          )}
          DialogProps={{ sx: styles.mobiledialogprops }}
          PaperProps={{ sx: styles.paperprops }}
          renderDay={renderWeekEndPickerDay}
        />
      </Box>
    </LocalizationProvider>
  )
}

export default CustomDatePicker
