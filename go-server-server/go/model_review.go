/*
 * Blog API
 *
 * This is a mini blog.
 *
 * API version: 1.0.0
 * Contact: chris-ju@qq.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Review struct {

	User *User `json:"user,omitempty"`

	Content string `json:"content,omitempty"`
}
