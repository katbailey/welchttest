package welchttest

import (
	. "github.com/ematvey/go-fn/fn"
	"math"
)

var sqrt func(float64) float64 = math.Sqrt
var pow func(float64, float64) float64 = math.Pow

// Cumulative Distribution Function for the t distrition.
//
// Returns the probability of getting a value greater than the passed in test
// statistic (t_score), with the specified degrees of freedom (dgf), if upper_tail
// is true. If upper_tail is false it returns the probability of getting a value
// less than or equal to the test statistic.
//
//                 * *
//               *     *
//              *       *
//             *         *
//            *           *    Upper tail
//           *            **    /
//          *             ***  /
//        *               *****
//  * * *                 ***********
// -----------------|-----|-----------------
//                  0    t-score
func StudentsT_CDF(t_score float64, dgf float64, upper_tail bool) float64 {
	if dgf < 1 {
		return float64(0)
	}
	// The cumulative distribution function for the Student's T distribution can be
	// expressed in terms of the regularized incomplete beta function. See
	// https://en.wikipedia.org/wiki/Student%27s_t-distribution#Cumulative_distribution_function
	inc_beta := BetaIncReg(dgf/2, 0.5, (dgf / ((t_score * t_score) + dgf)))
	// We use inc_beta / 2 if the t-score is negative and we're looking for the *lower*
	// tail or if the t-score is positive and we're looking for the *upper* tail. Otherwise
	// we return 1 - (inc_beta / 2).
	if (t_score < 0) != upper_tail {
		return inc_beta / 2
	} else {
		return 1 - (inc_beta / 2)
	}
}

// Returns the number of degress of freedom when comparing two sets of data with
// different variances.
func GetDegreesOfFreedom(nx int, ny int, varx float64, vary float64) float64 {
	if nx < 2 || ny < 2 {
		return float64(0)
	}
	denom := (pow(varx/float64(nx), 2) / (float64(nx) - 1)) + (pow(vary/float64(ny), 2) / float64(ny-1))
	if denom == 0 {
		return float64(0)
	}
	return pow((varx/float64(nx))+(vary/float64(ny)), 2) / denom
}

// Returns a test statistic for the difference in means between two sets of data
// with different variances.
func CalculateTScore(nx int, ny int, meanx float64, meany float64, varx float64, vary float64) float64 {
	var tscore float64
	if nx < 1 || ny < 1 {
		return tscore
	}
	denom := sqrt((varx / float64(nx)) + (vary / float64(ny)))
	if denom == 0 {
		return float64(0)
	}
	tscore = (meanx - meany) / denom
	return tscore
}

// Returns a value between 0 and 1 representing the degree of confidence that
// the true mean of x is greater than the true mean of y.
func GetConfidence(nx int, ny int, meanx float64, meany float64, varx float64, vary float64) float64 {
	var confidence float64
	if nx < 1 || ny < 1 || meanx < meany {
		return confidence
	}
	dgfree := GetDegreesOfFreedom(nx, ny, varx, vary)
	if dgfree == 0 {
		return confidence
	}
	var t_score float64
	t_score = CalculateTScore(nx, ny, meanx, meany, varx, vary)
	// If the t-score is negative or 0, then we have no confidence.
	if t_score <= 0 {
		return confidence
	}

	// The CDF function gives us either the probability of getting a result
	// less than or equal to our t-score, or the probability of getting a result
	// greater than our t-score, depending whether we set the "upper_tail" param
	// to true or false. Since we know our t-score is positive, we want the upper
	// tail.
	pvalue := StudentsT_CDF(t_score, dgfree, true)
	// No point in expressing negative confidence values so don't include it if our
	// p-value is greater than or equal to 0.5.
	if pvalue > 0.5 {
		return confidence
	}
	// For a two-sided test, the probability of as extreme a value or more extreme
	// needs to include both tails of the distribution, so we double the p-value
	// we got. Then 1 - pvalue is a reasonable representation of our confidence that
	// the true mean of x is greater than the true mean of y.
	confidence = 1 - (2 * pvalue)
	return confidence
}
