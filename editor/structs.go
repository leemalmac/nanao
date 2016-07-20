package editor

import "bytes"
import "../terminal"


type Editor interface {
  Open(path string)
  GetFilePath() string
  RefreshScreen()
  Edit()

  moveCursorUp()
  moveCursorDown()
  moveCursorLeft()
  moveCursorRight()
  boundCoursorRight()

  GetNumOfRows()
  ProcessKeyPress()
  getWingowSize()
}


type NanaoEditor struct {
  cursorXPos uint32 /* cursor x position */
  cursorYPos uint32 /* cursor y position */
  cursorXOffset int
  screenRows int32 /* Number of rows */
  screenCols int32 /* Number of columns */
  rowsOffset int32
  colsOffset int32
  isChanged bool /* Has a file been changed? */
  fileName string
  filePath string
  rows []Row /* File content */
  totalRowsNum int
  termOldState *terminal.State /* TODO: move it somewhere */
}


type Row struct {
  number uint32
  content *bytes.Buffer
  size int
}


type winsize struct {
  row    uint16
  col    uint16
  xpixel uint16
  ypixel uint16
}
