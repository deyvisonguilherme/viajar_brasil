(ns user
  (:require [mount.core :as mount]
            senem.core))

(defn start []
  (mount/start-without #'senem.core/http-server
                       #'senem.core/repl-server))

(defn stop []
  (mount/stop-except #'senem.core/http-server
                     #'senem.core/repl-server))

(defn restart []
  (stop)
  (start))


