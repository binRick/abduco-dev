package abduco

import "github.com/prometheus/procfs"

func getRelevantProcs(ppid int) (map[int]procfs.ProcStat, error) {
	fs, err := procfs.NewDefaultFS()
	if err != nil {
		return nil, err
	}
	procs := make(map[int]procfs.ProcStat)
	pending := make([]int, 1)
	pending[0] = ppid
	for len(pending) > 0 {
		curPid := pending[0]
		pending = pending[1:]
		cur, _ := fs.Proc(curPid)
		curStat, _ := cur.NewStat()
		procs[curPid] = curStat // Add to the set of procs we've found
		allProcs, _ := fs.AllProcs()
		for _, proc := range allProcs {
			stat, _ := proc.NewStat()
			//fmt.Printf("Compare %d %d/%d %d/%d\n",stat.PID, stat.PPID, curStat.PID, stat.PGRP, curStat.PGRP)
			if stat.PPID == curStat.PID {
				//fmt.Printf("%d\n", proc.PID)
				if _, ok := procs[stat.PID]; !ok {
					procs[stat.PID] = stat
					pending = append(pending, stat.PID)
				}
			}
		}
	}
	return procs, nil
}
