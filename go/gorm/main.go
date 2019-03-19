package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:l19880102./cloudboot?charset=utf8&parseTime=True&loc=Local")
}

func exec(db *gorm.DB) {
	var sql = `select t10.sn as sn,
	t10.job_id as job_id, 
	t5.oob->>'$.network.ip' as oob_ip,
	t10.start_time, 
	t10.end_time,
	t10.id, 
	t10.running_status, 
	t10.error, 
	t10.ipmi_result,
	t10.health_status,
	t10.created_at, 
	t10.updated_at from device t5 inner join inspection_latest t10 on t5.sn = t10.sn  
	where t10.id > 0 ? order by t10.end_time desc ?`

	var whereSQL strings.Builder
	if cond != nil {
		if cond.SN != "" {
			sns := strings.Split(cond.SN, ",")
			for i := range sns {
				sns[i] = fmt.Sprintf("'%s'", sns[i])
			}
			whereSQL.WriteString(fmt.Sprintf(" AND t5.sn IN(%s)", strings.Join(sns, ",")))
		}
		if cond.StartTime != "" {
			whereSQL.WriteString(fmt.Sprintf(" AND t10.start_time > '%s'", cond.StartTime))
		}
		if cond.EndTime != "" {
			whereSQL.WriteString(fmt.Sprintf(" AND t10.end_time < '%s'", cond.EndTime))
		}
		if cond.OOBIP != "" {
			whereSQL.WriteString(fmt.Sprintf(" AND t6.manage_ip = '%s'", cond.OOBIP))
		}
		if cond.RuningStatus != "" {
			whereSQL.WriteString(fmt.Sprintf(" AND t10.running_status = '%s'", cond.RuningStatus))
		}
		if cond.HealthStatus != "" {
			whereSQL.WriteString(fmt.Sprintf(" AND t10.health_status = '%s'", cond.HealthStatus))
		}
	}

	var limitSQL string
	if limiter != nil {
		limitSQL = fmt.Sprintf(" LIMIT %d,%d ", limiter.Offset, limiter.Limit)
	}
	// repo.log.Debugf("%s\n", fmt.Sprintf(sql, whereSQL.String(), limitSQL))
	if err := repo.db.Exec(sql, whereSQL.String(), limitSQL).Scan(&result).Error; err != nil {
		repo.log.Error(err)
		return nil, err
	}
}
