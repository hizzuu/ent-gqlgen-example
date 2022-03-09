package directive

import (
	"context"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/plate-backend/internal/domain"
	"github.com/pkg/errors"
)

func checkInt(v int, label string, min *int, max *int) error {
	if min != nil && v < *min {
		return domain.NewValidationError(errors.Errorf("%sは%d以上で入力してください。", label, *min))
	}

	if max != nil && v > *max {
		return domain.NewValidationError(errors.Errorf("%sは%d以下で入力してください。", label, *max))
	}

	return nil
}

func checkString(v string, label string, notEmpty *bool, notBlank *bool, pattern *string, min *int, max *int) error {
	if notEmpty != nil && v == "" {
		return domain.NewValidationError(errors.Errorf("%sを入力してください。", label))
	}

	if notBlank != nil {
		if strings.TrimSpace(v) == "" {
			return domain.NewValidationError(errors.Errorf("%sを入力してください。", label))
		}
	}

	if min != nil && utf8.RuneCountInString(v) < int(*min) {
		return domain.NewValidationError(errors.Errorf("%sは%d以上で入力してください。", label, *min))
	}

	if max != nil && utf8.RuneCountInString(v) > int(*max) {
		return domain.NewValidationError(errors.Errorf("%sは%d文字以下で入力してください。", label, *max))
	}

	if pattern != nil {
		if ok, _ := regexp.MatchString(*pattern, v); !ok {
			return domain.NewValidationError(errors.Errorf("%sのフォーマットが間違っています。", label))
		}
	}

	return nil
}

func (d *directive) Constraint(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
	label string,
	notEmpty *bool,
	notBlank *bool,
	pattern *string,
	min *int,
	max *int,
) (interface{}, error) {
	i, _ := next(ctx)
	switch v := i.(type) {
	case *int:
		if v != nil {
			if err := checkInt(*v, label, min, max); err != nil {
				return nil, err
			}
		}
	case int:
		if err := checkInt(v, label, min, max); err != nil {
			return nil, err
		}
	case *string:
		if v != nil {
			if err := checkString(*v, label, notEmpty, notBlank, pattern, min, max); err != nil {
				return nil, err
			}
		}
	case string:
		if err := checkString(v, label, notEmpty, notBlank, pattern, min, max); err != nil {
			return nil, err
		}
	}

	return i, nil
}
