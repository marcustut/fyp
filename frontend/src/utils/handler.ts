import { ErrorCode, FileRejection } from 'react-dropzone'
import { toast } from 'react-toastify'

export const handleFileRejections = (fileRejections: FileRejection[]) =>
  fileRejections.forEach((fileRejection) =>
    fileRejection.errors.forEach((fileError) => {
      switch (fileError.code) {
        case ErrorCode.FileInvalidType:
          toast.error(`${fileRejection.file.name} is not .PDF or .TXT`)
          break
        case ErrorCode.FileTooLarge:
          toast.error(`${fileRejection.file.name} is larger than 10MB`)
          break
        case ErrorCode.TooManyFiles:
          toast.error(`only 1 file is allowed at a time`)
          break
      }
    })
  )
