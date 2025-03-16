// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/templui/icons"
	"github.com/axzilla/templui/utils"
	"time"
)

type TimePickerProps struct {
	ID          string
	Name        string
	Value       time.Time
	Use12Hours  bool
	AMLabel     string
	PMLabel     string
	Placeholder string
	Required    bool
	Disabled    bool
	HasError    bool
	Class       string
	Attributes  templ.Attributes
}

func TimePickerScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 28, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('timePicker', () => ({\n\t\t\t\t\topen: false,\n\t\t\t\t\tformValue: null,\n\t\t\t\t\tdisplayValue: null,\n\t\t\t\t\tselectedHour: 0,\n\t\t\t\t\tselectedMinute: 0,\n\t\t\t\t\tisPM: false,\n\t\t\t\t\thours: [],\n\t\t\t\t\tminutes: [],\n\t\t\t\t\tuse12Hours: false,\n\t\t\t\t\tposition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tthis.use12Hours = this.$el.dataset.use12hours === 'true';\n\t\t\t\t\t\tif (this.$el.dataset?.value) {\n\t\t\t\t\t\t\tconst [hours, minutes] = this.$el.dataset.value.split(':').map(Number);\n\t\t\t\t\t\t\tthis.selectedMinute = minutes;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\t\tthis.isPM = hours >= 12;\n\t\t\t\t\t\t\t\tthis.selectedHour = hours > 12 ? hours - 12 : (hours === 0 ? 12 : hours);\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tthis.selectedHour = hours;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\ttoggleTimePicker() {\n\t\t\t\t\t\tthis.open = !this.open;\n\t\t\t\t\t\tif (this.open) {\n\t\t\t\t\t\t\tthis.$nextTick(() => this.updatePosition());\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tupdatePosition() {\n\t\t\t\t\t\tconst inputId = this.$root.dataset.inputId;\n\t\t\t\t\t\tconst trigger = document.getElementById(inputId);\n\t\t\t\t\t\tconst popup = this.$refs.timePickerPopup;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (!trigger || !popup) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst rect = trigger.getBoundingClientRect();\n\t\t\t\t\t\tconst popupRect = popup.getBoundingClientRect();\n\t\t\t\t\t\tconst viewportHeight = window.innerHeight;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (rect.bottom + popupRect.height > viewportHeight && rect.top > popupRect.height) {\n\t\t\t\t\t\t\tthis.position = 'top';\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.position = 'bottom';\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tpositionClass() {\n\t\t\t\t\t\treturn this.position === 'bottom' ? 'top-full mt-1' : 'bottom-full mb-1';\n\t\t\t\t\t},\n\n\t\t\t\t\tcloseTimePicker() {\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t},\n\n\t\t\t\t\tupdateValues() {\n\t\t\t\t\t\tlet hour24 = this.selectedHour;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tif (this.isPM && hour24 !== 12) hour24 += 12;\n\t\t\t\t\t\t\tif (!this.isPM && hour24 === 12) hour24 = 0;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.formValue = `${String(hour24).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')}`;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tthis.displayValue = `${String(this.selectedHour).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')} ${this.isPM ? 'PM' : 'AM'}`;\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.displayValue = this.formValue;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tselectHour() {\n\t\t\t\t\t\tconst hour = parseInt(this.$el.value);\n\t\t\t\t\t\tthis.selectedHour = hour;\n\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t},\n\n\t\t\t\t\tselectMinute() {\n\t\t\t\t\t\tconst minute = parseInt(this.$el.value);\n\t\t\t\t\t\tthis.selectedMinute = minute;\n\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t},\n\n\t\t\t\t\tselectPeriod() {\n\t\t\t\t\t\tconst period = this.$el.getAttribute('data-period');\n\t\t\t\t\t\tthis.isPM = period === 'PM';\n\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t},\n\n\t\t\t\t\tperiodButtonClass() {\n\t\t\t\t\t\tconst period = this.$el.getAttribute('data-period');\n\t\t\t\t\t\treturn period === 'PM' === this.isPM ? 'bg-primary text-primary-foreground' : '';\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func generateHourOptions(use12Hours bool) []SelectOption {
	options := make([]SelectOption, 0)
	if use12Hours {
		options = append(options, SelectOption{
			Label: "12",
			Value: "12",
		})
		for i := 1; i <= 11; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	} else {
		for i := 0; i <= 23; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	}
	return options
}

func generateMinuteOptions() []SelectOption {
	options := make([]SelectOption, 60)
	for i := 0; i < 60; i++ {
		options[i] = SelectOption{
			Label: fmt.Sprintf("%02d", i),
			Value: fmt.Sprintf("%d", i),
		}
	}
	return options
}

func TimePicker(props TimePickerProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if props.ID == "" {
			props.ID = utils.RandomID()
		}
		if props.Placeholder == "" {
			props.Placeholder = "Select time"
		}
		if props.AMLabel == "" {
			props.AMLabel = "AM"
		}
		if props.PMLabel == "" {
			props.PMLabel = "PM"
		}
		var templ_7745c5c3_Var5 = []any{utils.TwMerge("relative", props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Value != (time.Time{}) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value.Format("15:04"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 184, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, " data-use12hours=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%t", props.Use12Hours))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 186, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" x-data=\"timePicker\" data-input-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 188, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"><div class=\"relative\"><input type=\"hidden\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 193, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" x-bind:value=\"formValue\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Input(InputProps{
			ID:          props.ID,
			Value:       props.Value.Format("15:04"),
			Placeholder: props.Placeholder,
			Disabled:    props.Disabled,
			Class:       utils.TwMerge(props.Class, "peer"),
			HasError:    props.HasError,
			Type:        "text",
			Readonly:    true,
			Attributes: utils.MergeAttributes(
				templ.Attributes{
					"x-ref":  "timePickerInput",
					":value": "displayValue",
					"@click": "toggleTimePicker",
				},
				props.Attributes,
			),
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 = []any{
			utils.TwMerge(
				"absolute top-0 right-0 px-3 py-2",
				"cursor-pointer text-muted-foreground",
				"hover:text-foreground",
				"peer-disabled:pointer-events-none peer-disabled:opacity-50",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<button type=\"button\" @click=\"toggleTimePicker\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Disabled {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, " class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Clock(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</button></div><div x-show=\"open\" x-ref=\"timePickerPopup\" @click.away=\"closeTimePicker\" x-transition.opacity class=\"absolute left-0 z-50 w-72 p-4 rounded-lg border bg-popover shadow-md\" x-bind:class=\"positionClass\" @resize.window=\"updatePosition\"><div class=\"grid grid-cols-2 gap-2 flex-1\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Select(SelectProps{
			ID:   "hour-select",
			Name: "hour",
			Attributes: templ.Attributes{
				"x-bind:value": "selectedHour",
				"@change":      "selectHour",
			},
			Options: generateHourOptions(props.Use12Hours),
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Select(SelectProps{
			ID:   "minute-select",
			Name: "minute",
			Attributes: templ.Attributes{
				"x-bind:value": "selectedMinute",
				"@change":      "selectMinute",
			},
			Options: generateMinuteOptions(),
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</div></div><div class=\"flex justify-between mt-4 gap-2\"><div x-show=\"use12Hours\" class=\"flex justify-center gap-2\"><button type=\"button\" data-period=\"AM\" @click=\"selectPeriod\" class=\"px-3 py-1 text-sm rounded-md [&amp;:not(.bg-primary)]:hover:bg-muted\" x-bind:class=\"periodButtonClass\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(props.AMLabel)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 272, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</button> <button type=\"button\" data-period=\"PM\" @click=\"selectPeriod\" class=\"px-3 py-1 text-sm rounded-md [&amp;:not(.bg-primary)]:hover:bg-muted\" x-bind:class=\"periodButtonClass\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(props.PMLabel)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 281, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var15 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "Done")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Button(ButtonProps{
			Type:    "button",
			Size:    ButtonSizeSm,
			Variant: ButtonVariantSecondary,
			Attributes: templ.Attributes{
				"@click": "closeTimePicker",
			},
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var15), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
