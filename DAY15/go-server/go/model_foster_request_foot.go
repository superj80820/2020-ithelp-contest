/*
 * Digimon Service API
 *
 * 提供孵化數碼蛋與培育等數碼寶貝養成服務
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 培育所使用的食物
type FosterRequestFoot struct {
	// 食物名稱
	Name string `json:"name,omitempty"`
}
