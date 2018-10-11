module Complex where

import Control.Concurrent
import Control.Concurrent.Async
import Control.Monad
import Data.Foldable (traverse_)

complexMain :: IO ()
complexMain = do
  filePaths <- getFilePaths
  mapConcurrently_ readAndCountAllFiles filePaths

getFilePaths :: IO [FilePath]
getFilePaths = lines <$> getContents

printLinesOf :: String -> String -> IO ()
printLinesOf path contents =
  putStrLn $ path ++ ": " ++ show (length . lines $ contents)

readAndCountAllFiles :: String -> IO ()
readAndCountAllFiles path = do
  fileLines <- lines <$> readFile path
  mapConcurrently_ (\path -> readFile path >>= printLinesOf path) fileLines
