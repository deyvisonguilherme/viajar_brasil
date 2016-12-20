(ns senem.env
  (:require [clojure.tools.logging :as log]))

(def defaults
  {:init
   (fn []
     (log/info "\n-=[senem started successfully]=-"))
   :stop
   (fn []
     (log/info "\n-=[senem has shut down successfully]=-"))
   :middleware identity})
